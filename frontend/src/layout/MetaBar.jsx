import KpiCard from "../components/KpiCard";
import BeforeAfterCard from "../components/BeforeAfterCard";
import NotificationBell from "../components/NotificationBell";
import ServicePill from "../components/ServicePill";

function fmt(v) {
  return Number.isFinite(Number(v)) ? Number(v).toFixed(1) : "—";
}

export default function MetaBar({
  metrics,
  scenario,
  seed,
  services,
  gitops,
  onOpenWorkspace,
  highlightGraphNav = false,
}) {
  const prCount = Array.isArray(gitops?.prs) ? gitops.prs.length : 0;

  return (
    <div className="grid auto-rows-auto gap-3">
      <div className="grid grid-cols-[repeat(7,minmax(0,1fr))_220px] gap-3">
        <KpiCard label="Total Risk" value={fmt(metrics.totalRisk)} accent />
        <KpiCard label="Attack Paths" value={metrics.attackPathCount} onClick={() => onOpenWorkspace("attackpaths")} />
        <KpiCard label="Exposed Nodes" value={metrics.publicExposureCount} onClick={() => onOpenWorkspace("governance")} />
        <KpiCard label="Bill Shock" value={metrics.billShockCount} onClick={() => onOpenWorkspace("billshock")} />
        <KpiCard label="Compliance" value={fmt(metrics.complianceScore)} onClick={() => onOpenWorkspace("compliance")} />
        <KpiCard label="Cost Risk" value={fmt(metrics.costRiskScore)} />
        <KpiCard label="Blast Radius" value={metrics.maxBlastRadius} onClick={() => onOpenWorkspace("compliance")} />

        <div className="grid gap-2">
          <button
            onClick={() => onOpenWorkspace("graph")}
            className={`rounded-xl px-3 py-2 text-xs font-medium transition ${
              highlightGraphNav
                ? "border border-emerald-400/40 bg-emerald-500/15 text-emerald-200 shadow-[0_0_30px_rgba(16,185,129,0.25)]"
                : "border border-sky-500/20 bg-sky-500/10 text-sky-200"
            }`}
          >
            Open Graph Workspace
          </button>

          <button
            onClick={() => onOpenWorkspace("gitops")}
            className="rounded-xl border border-sky-500/20 bg-sky-500/10 px-3 py-2 text-xs font-medium text-sky-200"
          >
            Open GitOps Workspace
          </button>

          <button
            onClick={() => onOpenWorkspace("carbon")}
            className="rounded-xl border border-sky-500/20 bg-sky-500/10 px-3 py-2 text-xs font-medium text-sky-200"
          >
            Open Carbon Workspace
          </button>
        </div>
      </div>

      <div className="grid grid-cols-[1fr_1fr_1fr_1fr_320px] gap-3">
        <BeforeAfterCard
          title="Attack Paths"
          before={metrics.attackPathCount}
          after={metrics.projectedAttackPathCount}
          hint="Reachable attacker chains before and after remediation."
        />
        <BeforeAfterCard
          title="Compliance Score"
          before={fmt(metrics.complianceScore)}
          after={fmt(metrics.projectedComplianceScore)}
          hint="Current posture versus projected policy-safe posture."
        />
        <BeforeAfterCard
          title="Cost Risk"
          before={fmt(metrics.costRiskScore)}
          after={fmt(metrics.projectedCostRiskScore)}
          hint="FinOps pressure before and after governance actions."
        />
        <BeforeAfterCard
          title="Carbon"
          before={fmt(metrics.currentCarbonTotal)}
          after={fmt(metrics.projectedCarbonTotal)}
          hint="Backend-derived carbon intelligence before and after."
        />

        <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-3">
          <div className="flex items-start justify-between gap-3">
            <div className="min-w-0">
              <div className="text-[11px] uppercase tracking-wider text-slate-500">Run Context</div>
              <div className="mt-2 text-sm">
                Scenario: <span className="text-slate-200">{scenario || "-"}</span>
              </div>
              <div className="mt-1 text-sm">
                Seed: <span className="text-slate-200">{seed ?? "-"}</span>
              </div>
              <div className="mt-2 flex flex-wrap gap-2">
                {(services || []).slice(0, 4).map((svc) => (
                  <ServicePill key={svc.name} name={svc.name} status={svc.status} />
                ))}
              </div>
            </div>

            <NotificationBell
              count={prCount}
              onClick={() => onOpenWorkspace("gitops")}
            />
          </div>
        </div>
      </div>
    </div>
  );
}