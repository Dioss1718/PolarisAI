export default function BillShockTable({ items = [], onSelect }) {
  return (
    <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="text-lg font-semibold">Bill Shock Watchlist</div>

      <div className="mt-3 space-y-3">
        {items.map((f) => (
          <button
            key={f.nodeId}
            onClick={() => onSelect?.(f.nodeId)}
            className="block w-full rounded-xl border border-white/10 bg-slate-950/60 p-4 text-left hover:border-rose-500/30"
          >
            <div className="flex items-center justify-between">
              <div className="font-medium">{f.nodeId}</div>
              <div className={`text-xs ${f.billShock ? "text-rose-300" : "text-emerald-300"}`}>
                {f.billShock ? "SHOCK" : "STABLE"}
              </div>
            </div>

            <div className="mt-2 text-sm text-slate-300">
              Current Cost: {f.currentCost} · 30 Days: {f.forecast30} · 90 Days: {f.forecast90}
            </div>

            {f.shockReason ? (
              <div className="mt-2 text-xs text-slate-500">{f.shockReason}</div>
            ) : null}
          </button>
        ))}
      </div>
    </div>
  );
}