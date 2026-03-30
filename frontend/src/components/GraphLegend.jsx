export default function GraphLegend() {
  const items = [
    { label: "Critical Risk", color: "bg-rose-500" },
    { label: "High Risk", color: "bg-orange-500" },
    { label: "Medium Risk", color: "bg-yellow-400" },
    { label: "Low Risk", color: "bg-emerald-500" },
  ];

  return (
    <div className="flex flex-wrap items-center gap-2">
      <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-[11px] text-slate-300">
        AWS ◆
      </div>
      <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-[11px] text-slate-300">
        Azure ■
      </div>
      <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-[11px] text-slate-300">
        GCP ▲
      </div>

      {items.map((item) => (
        <div
          key={item.label}
          className="flex items-center gap-2 rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-[11px] text-slate-300"
        >
          <span className={`h-2.5 w-2.5 rounded-full ${item.color}`} />
          {item.label}
        </div>
      ))}

      <div className="rounded-full border border-rose-500/30 bg-rose-500/10 px-3 py-1 text-[11px] text-rose-200">
        Animated edge = selected attack path
      </div>
    </div>
  );
}