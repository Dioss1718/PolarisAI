export default function GraphActivityPanel({
  metrics,
  selectedAttackPath,
  onOpenWorkspace,
}) {
  const selectedPathLength = Array.isArray(selectedAttackPath) ? selectedAttackPath.length : 0;

  return (
    <div className="rounded-2xl border border-white/10 bg-white/[0.04] px-4 py-3">
      <div className="flex flex-wrap items-center justify-between gap-3">
        <div>
          <div className="text-[11px] uppercase tracking-[0.24em] text-slate-500">
            Unified Cloud Graph Activity
          </div>
          <div className="mt-1 text-sm text-slate-300">
            Live graph posture, attack-path visibility, and propagation-aware activity.
          </div>
        </div>

        <div className="flex flex-wrap gap-2">
          <MiniBadge label="Open Graph" value="Workspace" onClick={() => onOpenWorkspace?.("graph")} />
          <MiniBadge label="Attack Paths" value={metrics.attackPathCount} onClick={() => onOpenWorkspace?.("attackpaths")} />
          <MiniBadge label="Blast Radius" value={metrics.maxBlastRadius} onClick={() => onOpenWorkspace?.("compliance")} />
          <MiniBadge label="Selected Path" value={selectedPathLength || "None"} onClick={() => onOpenWorkspace?.("attackpaths")} />
          <MiniBadge label="Propagation" value="Open" onClick={() => onOpenWorkspace?.("propagation")} />
        </div>
      </div>
    </div>
  );
}

function MiniBadge({ label, value, onClick }) {
  return (
    <button
      onClick={onClick}
      className="rounded-full border border-white/10 bg-slate-950/70 px-3 py-1.5 text-xs text-slate-300 hover:border-sky-500/30"
    >
      <span className="text-slate-500">{label}:</span> {value}
    </button>
  );
}