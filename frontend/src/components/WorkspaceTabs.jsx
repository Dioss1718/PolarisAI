export default function WorkspaceTabs({ activeTab, setActiveTab }) {
  const tabs = [
    ["graph", "Graph Workspace"],
    ["attackpaths", "Attack Paths"],
    ["governance", "Governance Actions"],
    ["explainability", "Explainability"],
    ["billshock", "Bill Shock Watch"],
    ["gitops", "GitOps"],
    ["feedback", "Adaptive Feedback"],
    ["carbon", "Carbon Intelligence"],
    ["negotiation", "Negotiation & Tradeoffs"],
    ["compliance", "Compliance & Blast Radius"],
    ["propagation", "Risk Propagation"],
  ];

  return (
    <div className="flex flex-nowrap items-center gap-2 overflow-x-auto whitespace-nowrap">
      {tabs.map(([key, label]) => (
        <button
          key={key}
          onClick={() => setActiveTab(key)}
          className={`shrink-0 rounded-lg border px-3 py-2 text-xs ${
            activeTab === key
              ? "border-sky-500/30 bg-sky-500/10 text-sky-200"
              : "border-white/10 bg-slate-950/60 text-slate-300"
          }`}
        >
          {label}
        </button>
      ))}
    </div>
  );
}