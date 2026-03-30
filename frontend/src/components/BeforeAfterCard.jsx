export default function BeforeAfterCard({ title, before, after, hint }) {
  return (
    <div className="flex h-full min-w-0 flex-col justify-center rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="truncate text-[11px] uppercase tracking-[0.25em] text-slate-500">
        {title}
      </div>

      <div className="mt-3 grid grid-cols-2 gap-3">
        <div className="min-w-0">
          <div className="text-xs text-slate-500">Before</div>
          <div className="mt-1 truncate text-xl font-semibold text-rose-300">{before}</div>
        </div>
        <div className="min-w-0">
          <div className="text-xs text-slate-500">After</div>
          <div className="mt-1 truncate text-xl font-semibold text-emerald-300">{after}</div>
        </div>
      </div>

      <div className="mt-3 line-clamp-2 text-xs text-slate-500">{hint}</div>
    </div>
  );
}