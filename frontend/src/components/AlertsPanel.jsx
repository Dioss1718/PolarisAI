export default function AlertsPanel({ alerts = [], onOpenWorkspace }) {
  if (!alerts.length) return null;

  return (
    <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-3">
      <div className="flex flex-wrap gap-3">
        {alerts.map((alert, index) => (
          <button
            key={`${alert.title}-${index}`}
            onClick={() => onOpenWorkspace?.(alert.workspaceTab || "governance")}
            className={`rounded-xl border px-3 py-2 text-left text-sm transition hover:scale-[1.01] ${
              alert.severity === "critical"
                ? "border-rose-500/30 bg-rose-500/10 text-rose-200"
                : alert.severity === "high"
                ? "border-amber-500/30 bg-amber-500/10 text-amber-200"
                : alert.severity === "medium"
                ? "border-sky-500/30 bg-sky-500/10 text-sky-200"
                : "border-emerald-500/30 bg-emerald-500/10 text-emerald-200"
            }`}
          >
            <div className="font-medium">{alert.title}</div>
            <div className="text-xs opacity-90">{alert.metric}</div>
            <div className="mt-1 text-[11px] opacity-80">{alert.reason}</div>
          </button>
        ))}
      </div>
    </div>
  );
}