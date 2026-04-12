import { useEffect, useMemo, useState } from "react";

const scenarioOptions = [
  "FULL_CHAOS",
  "SECURITY_BREACH",
  "COST_SPIKE",
  "POLICY_DRIFT",
];

function deltaTone(value, inverse = false) {
  const n = Number(value || 0);
  if (n === 0) return "text-slate-300";
  const good = inverse ? n < 0 : n > 0;
  return good ? "text-emerald-300" : "text-rose-300";
}

function formatDelta(value) {
  const n = Number(value || 0);
  if (n > 0) return `+${n}`;
  return `${n}`;
}

function buildEditableManualStateFromCurrentState(currentState) {
  if (!currentState || !Array.isArray(currentState.nodes) || !Array.isArray(currentState.edges)) {
    return { nodes: [], edges: [] };
  }

  return {
    nodes: currentState.nodes.map((node) => ({
      id: node.id,
      type: node.type,
      name: node.label || node.id,
      cloud_provider: node.cloud,
      region: node.region,
      environment: node.environment,
      cost: Number(node.cost ?? 0),
      utilization: Number(node.utilization ?? 0),
      exposure: node.exposure,
      criticality: Number(node.criticality ?? 1),
      compliance_flags: Array.isArray(node.compliance) ? node.compliance : [],
    })),
    edges: currentState.edges.map((edge) => ({
      from: edge.from,
      to: edge.to,
      type: edge.type,
      weight: Number(edge.weight ?? 1),
    })),
  };
}

function safeJSONStringify(value) {
  try {
    return JSON.stringify(value, null, 2);
  } catch {
    return JSON.stringify({ nodes: [], edges: [] }, null, 2);
  }
}

function buildInitialManualJSON(currentState) {
  return safeJSONStringify(buildEditableManualStateFromCurrentState(currentState));
}

export default function SimulationStudio({
  state,
  onCompare,
  loading,
  allowed = true,
  scenario = "FULL_CHAOS",
  seed = 42,
}) {
  const [manualJSON, setManualJSON] = useState(buildInitialManualJSON(state));
  const [baseScenario, setBaseScenario] = useState(scenario || "FULL_CHAOS");
  const [candidateScenario, setCandidateScenario] = useState("POLICY_DRIFT");
  const [candidateSeed, setCandidateSeed] = useState(seed || 42);
  const [compareMode, setCompareMode] = useState("manual");
  const [compareResult, setCompareResult] = useState(null);
  const [compareError, setCompareError] = useState("");
  const [reloadMessage, setReloadMessage] = useState("");

  useEffect(() => {
    setBaseScenario(scenario || "FULL_CHAOS");
  }, [scenario]);

  useEffect(() => {
    setCandidateSeed(seed || 42);
  }, [seed]);

  useEffect(() => {
    setManualJSON(buildInitialManualJSON(state));
  }, [state]);

  const baselineLabel = useMemo(
    () => `${baseScenario} · seed ${seed || 42}`,
    [baseScenario, seed]
  );

  const loadCurrentStateIntoEditor = () => {
    const updated = buildInitialManualJSON(state);
    setManualJSON(updated);
    setReloadMessage(`Loaded current dashboard state at ${new Date().toLocaleTimeString()}`);
  };

  const submitCompare = async () => {
    setCompareError("");
    setCompareResult(null);

    try {
      let payload = {
        baseScenario,
        baseSeed: seed || 42,
        candidateScenario,
        candidateSeed,
      };

      if (compareMode === "manual") {
        const parsed = JSON.parse(manualJSON);
        payload = {
          baseScenario,
          baseSeed: seed || 42,
          candidateSeed,
          manualData: parsed,
        };
      }

      if (!onCompare) {
        throw new Error("SimulationStudio is missing onCompare handler");
      }

      const result = await onCompare(payload);

      if (!result || !result.delta || !result.baseline || !result.candidate) {
        throw new Error("Compare API returned incomplete result");
      }

      setCompareResult(result);
    } catch (err) {
      setCompareError(
        err?.response?.data?.error ||
          err?.message ||
          "What-if comparison failed"
      );
    }
  };

  if (!allowed) {
    return (
      <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <div className="text-lg font-semibold">Simulation Studio</div>
        <div className="mt-2 text-sm text-slate-400">
          Your role does not have access to simulation studio.
        </div>
      </div>
    );
  }

  const delta = compareResult?.delta || {};

  return (
    <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="flex flex-wrap items-start justify-between gap-3">
        <div>
          <div className="text-lg font-semibold">Simulation Studio</div>
          <p className="mt-1 text-sm text-slate-400">
            Edit the current scenario state and compare it against the baseline.
          </p>
        </div>

        <button
          onClick={loadCurrentStateIntoEditor}
          disabled={loading}
          className="rounded-xl border border-white/10 bg-slate-900 px-3 py-2 text-sm text-slate-300 disabled:opacity-50"
        >
          Reload Current State
        </button>
      </div>

      <div className="mt-4 grid gap-3 lg:grid-cols-4">
        <div className="rounded-xl bg-slate-900/70 p-3 lg:col-span-2">
          <div className="text-[11px] uppercase tracking-wider text-slate-500">
            Baseline Source
          </div>
          <div className="mt-1 text-sm text-slate-200">
            {scenario || "FULL_CHAOS"} · seed {seed || 42}
          </div>
        </div>

        <label className="rounded-xl bg-slate-900/70 p-3">
          <div className="text-[11px] uppercase tracking-wider text-slate-500">
            Comparison Mode
          </div>
          <select
            value={compareMode}
            onChange={(e) => setCompareMode(e.target.value)}
            className="mt-2 w-full rounded-xl border border-white/10 bg-slate-950 px-3 py-2 text-sm outline-none"
          >
            <option value="manual">Manual vs Baseline</option>
            <option value="scenario">Scenario vs Baseline</option>
          </select>
        </label>

        <label className="rounded-xl bg-slate-900/70 p-3">
          <div className="text-[11px] uppercase tracking-wider text-slate-500">
            Candidate Seed
          </div>
          <input
            type="number"
            value={candidateSeed}
            onChange={(e) => setCandidateSeed(Number(e.target.value || 42))}
            className="mt-2 w-full rounded-xl border border-white/10 bg-slate-950 px-3 py-2 text-sm outline-none"
          />
        </label>
      </div>

      {compareMode === "scenario" ? (
        <div className="mt-3 grid gap-3 md:grid-cols-2">
          <label className="rounded-xl bg-slate-900/70 p-3">
            <div className="text-[11px] uppercase tracking-wider text-slate-500">
              Baseline Scenario
            </div>
            <select
              value={baseScenario}
              onChange={(e) => setBaseScenario(e.target.value)}
              className="mt-2 w-full rounded-xl border border-white/10 bg-slate-950 px-3 py-2 text-sm outline-none"
            >
              {scenarioOptions.map((option) => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </select>
          </label>

          <label className="rounded-xl bg-slate-900/70 p-3">
            <div className="text-[11px] uppercase tracking-wider text-slate-500">
              Candidate Scenario
            </div>
            <select
              value={candidateScenario}
              onChange={(e) => setCandidateScenario(e.target.value)}
              className="mt-2 w-full rounded-xl border border-white/10 bg-slate-950 px-3 py-2 text-sm outline-none"
            >
              {scenarioOptions.map((option) => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </select>
          </label>
        </div>
      ) : null}

      {reloadMessage ? (
        <div className="mt-3 rounded-xl border border-emerald-500/20 bg-emerald-500/10 p-3 text-sm text-emerald-300">
          {reloadMessage}
        </div>
      ) : null}

      <div className="mt-4">
        <div className="mb-2 text-sm font-medium text-slate-200">
          Editable Candidate JSON
        </div>
        <textarea
          value={manualJSON}
          onChange={(e) => setManualJSON(e.target.value)}
          className="h-[380px] w-full rounded-xl border border-white/10 bg-slate-950/80 p-4 text-xs outline-none"
        />
      </div>

     

      <div className="mt-4 flex flex-wrap gap-3">
        <button
          onClick={submitCompare}
          disabled={loading}
          className="rounded-xl border border-emerald-500/30 bg-emerald-500/10 px-4 py-2 text-sm text-emerald-300 disabled:opacity-50"
        >
          Run What-If Comparison
        </button>
      </div>

      {compareError ? (
        <div className="mt-4 rounded-xl border border-rose-500/30 bg-rose-500/10 p-3 text-sm text-rose-200">
          {compareError}
        </div>
      ) : null}

      {compareResult ? (
        <div className="mt-6 space-y-4">
          <div className="rounded-xl bg-slate-900/70 p-3 text-sm text-slate-300">
            <div className="font-medium text-slate-200">Comparison Summary</div>
            <div className="mt-1">Mode: {compareResult.mode}</div>
            <div className="mt-1 text-xs text-slate-500">
              Baseline: {baselineLabel}
            </div>
          </div>

          <div className="grid gap-3 md:grid-cols-2 xl:grid-cols-5">
            <MetricCard label="Cost Δ" value={formatDelta(delta.currentTotalCostDelta)} tone={deltaTone(delta.currentTotalCostDelta, true)} />
            <MetricCard label="Forecast 30 Δ" value={formatDelta(delta.forecast30TotalDelta)} tone={deltaTone(delta.forecast30TotalDelta, true)} />
            <MetricCard label="Forecast 90 Δ" value={formatDelta(delta.forecast90TotalDelta)} tone={deltaTone(delta.forecast90TotalDelta, true)} />
            <MetricCard label="Attack Paths Δ" value={formatDelta(delta.attackPathCountDelta)} tone={deltaTone(delta.attackPathCountDelta, true)} />
            <MetricCard label="Public Exposure Δ" value={formatDelta(delta.publicExposureCountDelta)} tone={deltaTone(delta.publicExposureCountDelta, true)} />
            <MetricCard label="High Risk Nodes Δ" value={formatDelta(delta.highRiskCountDelta)} tone={deltaTone(delta.highRiskCountDelta, true)} />
            <MetricCard label="Bill Shock Δ" value={formatDelta(delta.billShockCountDelta)} tone={deltaTone(delta.billShockCountDelta, true)} />
            <MetricCard label="Average Risk Δ" value={formatDelta(delta.averageRiskDelta)} tone={deltaTone(delta.averageRiskDelta, true)} />
            <MetricCard label="Compliance Δ" value={formatDelta(delta.complianceScoreDelta)} tone={deltaTone(delta.complianceScoreDelta, false)} />
            <MetricCard label="Cost Risk Δ" value={formatDelta(delta.costRiskScoreDelta)} tone={deltaTone(delta.costRiskScoreDelta, true)} />
          </div>

          <div className="grid gap-3 md:grid-cols-2">
            <RunSnapshot title="Baseline" run={compareResult.baseline} />
            <RunSnapshot title="Candidate" run={compareResult.candidate} />
          </div>
        </div>
      ) : null}
    </div>
  );
}

function MetricCard({ label, value, tone }) {
  return (
    <div className="rounded-xl bg-slate-900/70 p-3">
      <div className="text-[11px] uppercase tracking-wider text-slate-500">{label}</div>
      <div className={`mt-1 text-lg font-semibold ${tone}`}>{value}</div>
    </div>
  );
}

function RunSnapshot({ title, run }) {
  const summary = run?.summary || {};
  return (
    <div className="rounded-xl border border-white/10 bg-slate-900/70 p-3">
      <div className="font-semibold text-slate-200">{title}</div>
      <div className="mt-3 grid gap-2 text-sm text-slate-300">
        <div>Total Cost: {summary.currentTotalCost ?? "—"}</div>
        <div>Forecast 30: {summary.forecast30Total ?? "—"}</div>
        <div>Forecast 90: {summary.forecast90Total ?? "—"}</div>
        <div>Attack Paths: {summary.attackPathCount ?? "—"}</div>
        <div>Public Exposure: {summary.publicExposureCount ?? "—"}</div>
        <div>High Risk Nodes: {summary.highRiskCount ?? "—"}</div>
        <div>Compliance: {summary.complianceScore ?? "—"}</div>
        <div>Average Risk: {summary.averageRisk ?? "—"}</div>
      </div>
    </div>
  );
}