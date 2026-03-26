export default function ForecastPanel({ forecasts, onSelect }) {
  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Bill Shock & Cost Trajectory</h2>
      <div className="mt-4 grid gap-3 md:grid-cols-2 xl:grid-cols-3">
        {forecasts.map((f) => (
          <button
            key={f.nodeId}
            onClick={() => onSelect(f.nodeId)}
            className="rounded-xl border border-borderSoft bg-slate-950/70 p-4 text-left hover:border-emerald-500/40"
          >
            <div className="flex items-center justify-between">
              <div className="font-medium">{f.nodeId}</div>
              <div className={`text-xs ${f.billShock ? "text-rose-300" : "text-emerald-300"}`}>
                {f.billShock ? "SHOCK" : "STABLE"}
              </div>
            </div>
            <div className="mt-3 text-sm text-slate-300">
              <div>Current: {f.currentCost}</div>
              <div>30d: {f.forecast30}</div>
              <div>90d: {f.forecast90}</div>
            </div>
            {f.shockReason ? (
              <div className="mt-3 text-xs text-slate-500">{f.shockReason}</div>
            ) : null}
          </button>
        ))}
      </div>
    </div>
  );
}