export default function DecisionRail({ stages }) {
  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Execution Rail</h2>
      <div className="mt-4 space-y-3">
        {stages.map((stage) => (
          <div
            key={stage.name}
            className="flex items-center justify-between rounded-xl border border-borderSoft bg-slate-950/60 px-4 py-3"
          >
            <div className="text-sm font-medium">{stage.name}</div>
            <div
              className={`rounded-full px-3 py-1 text-xs uppercase tracking-wider ${
                stage.status === "complete"
                  ? "bg-emerald-500/15 text-emerald-300"
                  : stage.status === "ready"
                  ? "bg-sky-500/15 text-sky-300"
                  : "bg-amber-500/15 text-amber-300"
              }`}
            >
              {stage.status}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}