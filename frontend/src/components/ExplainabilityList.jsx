function cleanExplanation(text = "") {
  return String(text)
    .replace(/\*\*/g, "")
    .replace(/\*/g, "")
    .replace(/#+/g, "")
    .replace(/\s+/g, " ")
    .trim();
}

export default function ExplainabilityList({ explanations = [], onSelect }) {
  return (
    <div className="space-y-4">
      {explanations.map((item, index) => (
        <button
          key={`${item.nodeId}-${index}`}
          onClick={() => onSelect?.(item.nodeId)}
          className="block w-full rounded-2xl border border-white/10 bg-slate-950/60 p-4 text-left hover:border-sky-500/30"
        >
          <div className="flex items-center justify-between gap-3">
            <div>
              <div className="text-lg font-semibold">{item.nodeId}</div>
              <div className="text-sm text-slate-500">{item.action}</div>
            </div>
          </div>

          <div className="mt-3 rounded-xl bg-slate-900/70 p-4 text-sm leading-7 text-slate-300">
            {cleanExplanation(item.explanation)}
          </div>
        </button>
      ))}
    </div>
  );
}