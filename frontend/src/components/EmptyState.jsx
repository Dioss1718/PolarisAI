export default function EmptyState({ title, subtitle }) {
  return (
    <div className="flex h-full items-center justify-center rounded-2xl border border-white/10 bg-white/[0.04] p-6 text-center">
      <div>
        <div className="text-lg font-semibold">{title}</div>
        <div className="mt-2 text-sm text-slate-400">{subtitle}</div>
      </div>
    </div>
  );
}