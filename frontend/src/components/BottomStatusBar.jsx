export default function BottomStatusBar({
  session,
  services = [],
  metrics,
  autoRefreshOn,
}) {
  const role = session?.role || "GUEST";
  const servicesUp = Array.isArray(services)
    ? services.filter((s) => String(s.status || "").toUpperCase() === "UP").length
    : 0;

  return (
    <div className="fixed bottom-0 left-0 right-0 z-[80] border-t border-white/10 bg-slate-950/95 backdrop-blur">
      <div className="mx-auto flex min-h-[44px] w-full max-w-[1920px] flex-wrap items-center gap-x-5 gap-y-1 px-4 py-2 text-xs text-slate-300">
        <span>Role: {role}</span>
        <span>Total Risk: {Number(metrics?.totalRisk || 0).toFixed(1)}</span>
        <span>Compliance: {Number(metrics?.compliance || 0).toFixed(1)}</span>
        <span>90-day Forecast: {Number(metrics?.forecast90 || 0).toFixed(1)}</span>
        <span>Green Score: {Number(metrics?.greenScore || 0).toFixed(1)}</span>
        <span>Services Up: {servicesUp}</span>
        <span className="font-medium text-emerald-300">
          Projected Improvement: {Number(metrics?.projectedImprovement || 0).toFixed(1)}%
        </span>
        <span className="ml-auto text-slate-400">
          Auto-refresh: {autoRefreshOn ? "ON" : "OFF"}
        </span>
      </div>
    </div>
  );
}