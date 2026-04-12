function cleanExplanation(text = "") {
  return String(text)
    .replace(/\*\*/g, "")
    .replace(/\*/g, "")
    .replace(/#+/g, "")
    .replace(/\s+/g, " ")
    .trim();
}

function formatSourceLabel(source = "") {
  const value = String(source || "").trim();
  if (!value) return "Unknown source";
  return value.replace(/_/g, " ");
}

export default function ExplainabilityList({ explanations = [], onSelect }) {
  return (
    <div className="space-y-4">
      {explanations.map((item, index) => {
        const grounded = Boolean(item.grounded);
        const sources = Array.isArray(item.sources) ? item.sources : [];

        return (
          <button
            key={`${item.nodeId}-${index}`}
            onClick={() => onSelect?.(item.nodeId)}
            className="block w-full rounded-2xl border border-white/10 bg-slate-950/60 p-4 text-left hover:border-sky-500/30"
          >
            <div className="flex items-start justify-between gap-3">
              <div>
                <div className="text-lg font-semibold">{item.nodeId}</div>
                <div className="text-sm text-slate-500">{item.action}</div>
              </div>

              <div
                className={`rounded-full px-3 py-1 text-xs font-medium ${
                  grounded
                    ? "border border-emerald-500/30 bg-emerald-500/10 text-emerald-300"
                    : "border border-amber-500/30 bg-amber-500/10 text-amber-300"
                }`}
              >
                {grounded ? "Grounded" : "Ungrounded"}
              </div>
            </div>

            <div className="mt-3 rounded-xl bg-slate-900/70 p-4 text-sm leading-7 text-slate-300">
              {cleanExplanation(item.explanation)}
            </div>

            <div className="mt-4">
              <div className="text-xs font-semibold uppercase tracking-wide text-slate-400">
                Evidence Sources
              </div>

              {sources.length > 0 ? (
                <div className="mt-2 flex flex-wrap gap-2">
                  {sources.map((source, sourceIndex) => (
                    <span
                      key={`${item.nodeId}-source-${sourceIndex}`}
                      className="rounded-full border border-sky-500/20 bg-sky-500/10 px-3 py-1 text-xs text-sky-200"
                    >
                      {formatSourceLabel(source)}
                    </span>
                  ))}
                </div>
              ) : (
                <div className="mt-2 text-sm text-slate-500">
                  No evidence sources were returned for this explanation.
                </div>
              )}
            </div>
          </button>
        );
      })}
    </div>
  );
}