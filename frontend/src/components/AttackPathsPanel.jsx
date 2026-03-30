import EmptyState from "./EmptyState";

export default function AttackPathsPanel({ attackPaths = [], onSelectPath, selectedIndex }) {
  if (!attackPaths.length) {
    return (
      <EmptyState
        title="No attack paths discovered"
        subtitle="Run governance or load a scenario to inspect reachable attack paths."
      />
    );
  }

  return (
    <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="flex items-center justify-between">
        <div>
          <div className="text-lg font-semibold">Attack Paths</div>
          <div className="text-sm text-slate-400">
            Real paths discovered from public entry points to sensitive targets.
          </div>
        </div>
        <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-xs">
          {attackPaths.length} paths
        </div>
      </div>

      <div className="mt-4 space-y-3">
        {attackPaths.map((path, index) => (
          <button
            key={`attack-path-${index}`}
            onClick={() => onSelectPath(index)}
            className={`block w-full rounded-xl border p-3 text-left ${
              selectedIndex === index
                ? "border-rose-400/40 bg-rose-500/10"
                : "border-white/10 bg-slate-950/60"
            }`}
          >
            <div className="text-xs uppercase tracking-wider text-slate-500">Path {index + 1}</div>
            <div className="mt-2 text-sm text-slate-200 break-words">
              {path.join(" → ")}
            </div>
          </button>
        ))}
      </div>
    </div>
  );
}