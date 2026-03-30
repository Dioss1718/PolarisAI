export default function KpiCard({ label, value, accent = false, onClick }) {
  return (
    <button
      onClick={onClick}
      className={`flex h-full min-w-0 flex-col justify-center rounded-2xl border px-3 py-3 text-left transition hover:scale-[1.01] ${
        accent
          ? "border-rose-500/25 bg-rose-500/10"
          : "border-white/10 bg-white/[0.04]"
      }`}
    >
      <div className="truncate text-[10px] uppercase tracking-[0.22em] text-slate-500">
        {label}
      </div>
      <div className="mt-2 truncate text-2xl font-semibold">{value}</div>
    </button>
  );
}