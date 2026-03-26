export default function ExplainabilityPanel({ explanations, onSelect }) {
  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Explainability Feed</h2>
      <div className="mt-4 space-y-3 max-h-[420px] overflow-auto">
        {explanations.map((item) => (
          <button
            key={`${item.nodeId}-${item.action}`}
            onClick={() => onSelect(item.nodeId)}
            className="w-full rounded-xl border border-borderSoft bg-slate-950/70 p-4 text-left hover:border-violet-500/40"
          >
            <div className="text-sm font-medium text-violet-300">{item.nodeId} · {item.action}</div>
            <div className="mt-2 line-clamp-4 text-sm text-slate-300">{item.explanation}</div>
          </button>
        ))}
      </div>
    </div>
  );
}