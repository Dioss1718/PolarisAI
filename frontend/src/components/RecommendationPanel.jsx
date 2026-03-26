export default function RecommendationPanel({ recommendations, onSelect }) {
  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Governance Decisions</h2>
      <p className="mt-1 text-sm text-slate-400">
        Final policy-validated actions selected by the negotiation engine.
      </p>

      <div className="mt-4 space-y-3 max-h-[420px] overflow-auto">
        {recommendations.map((r) => (
          <button
            key={`${r.nodeId}-${r.finalAction}`}
            onClick={() => onSelect(r)}
            className="w-full rounded-xl border border-borderSoft bg-slate-950/70 p-4 text-left transition hover:border-sky-500/40 hover:bg-slate-900"
          >
            <div className="flex items-center justify-between">
              <div>
                <div className="font-medium">{r.nodeId}</div>
                <div className="text-xs text-slate-500">{r.cloud} · {r.type} · {r.environment}</div>
              </div>
              <div className="text-right">
                <div className="text-sm text-sky-300">{r.finalAction}</div>
                <div className="text-xs text-slate-500">{r.status} · {r.score.toFixed(2)}</div>
              </div>
            </div>
            <div className="mt-3 text-sm text-slate-300">{r.reason}</div>
          </button>
        ))}
      </div>
    </div>
  );
}