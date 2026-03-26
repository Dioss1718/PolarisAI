import { Activity, Play, RefreshCcw, Shield, GitBranch } from "lucide-react";

export default function TopBar({
  onRun,
  loading,
  scenario,
  setScenario,
  seed,
  setSeed,
}) {
  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow backdrop-blur p-4">
      <div className="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
        <div>
          <div className="text-xs uppercase tracking-[0.3em] text-sky-300/70">
            Autonomous Cloud Governance Engine
          </div>
          <h1 className="mt-1 text-3xl font-semibold tracking-tight">
            Polaris Control Plane
          </h1>
          <p className="mt-1 text-sm text-slate-400">
            Multi-cloud graph intelligence, safe remediation, policy enforcement, AI explainability, and GitOps execution.
          </p>
        </div>

        <div className="grid grid-cols-1 gap-3 sm:grid-cols-4">
          <div className="rounded-xl border border-borderSoft bg-slate-950/60 px-4 py-3">
            <div className="flex items-center gap-2 text-slate-400 text-xs uppercase">
              <Activity size={14} />
              Scenario
            </div>
            <select
              value={scenario}
              onChange={(e) => setScenario(e.target.value)}
              className="mt-2 w-full rounded-lg border border-borderSoft bg-slate-900 px-3 py-2 text-sm outline-none"
            >
              <option value="FULL_CHAOS">FULL_CHAOS</option>
              <option value="SECURITY_BREACH">SECURITY_BREACH</option>
              <option value="COST_SPIKE">COST_SPIKE</option>
              <option value="COMPLIANCE_DRIFT">COMPLIANCE_DRIFT</option>
              <option value="BILL_SHOCK">BILL_SHOCK</option>
            </select>
          </div>

          <div className="rounded-xl border border-borderSoft bg-slate-950/60 px-4 py-3">
            <div className="flex items-center gap-2 text-slate-400 text-xs uppercase">
              <Shield size={14} />
              Seed
            </div>
            <input
              value={seed}
              onChange={(e) => setSeed(Number(e.target.value || 0))}
              className="mt-2 w-full rounded-lg border border-borderSoft bg-slate-900 px-3 py-2 text-sm outline-none"
            />
          </div>

          <div className="rounded-xl border border-borderSoft bg-slate-950/60 px-4 py-3">
            <div className="flex items-center gap-2 text-slate-400 text-xs uppercase">
              <GitBranch size={14} />
              Mode
            </div>
            <div className="mt-2 text-sm text-slate-200">
              Live Operator Console
            </div>
            <div className="text-xs text-slate-500">
              Graph-aware governance
            </div>
          </div>

          <div className="flex items-end gap-2">
            <button
              onClick={onRun}
              disabled={loading}
              className="flex-1 rounded-xl bg-sky-500 px-4 py-3 font-medium text-slate-950 transition hover:bg-sky-400 disabled:opacity-50"
            >
              <div className="flex items-center justify-center gap-2">
                {loading ? <RefreshCcw size={16} className="animate-spin" /> : <Play size={16} />}
                {loading ? "Running..." : "Run Governance"}
              </div>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}