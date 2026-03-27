import { useState } from "react";

export default function ExplainabilityPanel({ explanations, onSelect }) {

  const [expanded, setExpanded] = useState(null);

  
  const cleanText = (text) => {
    if (!text) return "";
    return text
      .replace(/\*\*/g, "")
      .replace(/SUMMARY:/gi, "Summary:")
      .replace(/RISK_REASON:/gi, "Risk Reason:")
      .replace(/_/g, " ")
      .trim();
  };

  const toggleExpand = (key) => {
    setExpanded(expanded === key ? null : key);
  };

  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Explainability Feed</h2>

      <div className="mt-4 space-y-3 max-h-[420px] overflow-auto">
        {explanations.map((item) => {
          const key = `${item.nodeId}-${item.action}`;
          const isExpanded = expanded === key;
          const text = cleanText(item.explanation);

          return (
            <div
              key={key}
              className="w-full rounded-xl border border-borderSoft bg-slate-950/70 p-4"
            >
              <div
                onClick={() => onSelect(item.nodeId)}
                className="cursor-pointer text-sm font-medium text-violet-300"
              >
                {item.nodeId} · {item.action}
              </div>

              
              <div
                className={`mt-2 text-sm text-slate-300 whitespace-pre-wrap ${
                  isExpanded ? "" : "line-clamp-4"
                }`}
              >
                {text}
              </div>

              
              {text.length > 150 && (
                <button
                  onClick={() => toggleExpand(key)}
                  className="mt-2 text-xs text-sky-400 hover:text-sky-300"
                >
                  {isExpanded ? "Show less" : "Show more"}
                </button>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
}