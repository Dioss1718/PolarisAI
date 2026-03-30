export default function BottomStatusBar({ session, services, metrics, autoRefreshOn }) {
  const servicesUp = Array.isArray(services)
    ? services.filter((s) => s.status === "up").length
    : 0;
  const projectedImprovement = Number(metrics.projectedRiskReductionPct || 0);

  return (
    <div className="fixed bottom-0 left-0 right-0 z-40 border-t border-white/10 bg-slate-950/95 px-4 py-2 backdrop-blur-xl">
      <div className="mx-auto flex max-w-[1800px] flex-wrap items-center justify-between gap-3 text-[11px]">
        <div className="flex flex-wrap items-center gap-4 text-slate-400">
          <span>Role: {session?.role || "GUEST"}</span>
          <span>Total Risk: {Number(metrics.totalRisk || 0).toFixed(1)}</span>
          <span>Compliance: {Number(metrics.complianceScore || 0).toFixed(1)}</span>
          <span>90-day Forecast: {Number(metrics.forecast90Total || 0).toFixed(1)}</span>
          <span>Green Score: {Number(metrics.greenScore || 0).toFixed(1)}</span>
          <span>Services Up: {servicesUp}</span>
          <span className="font-semibold text-emerald-300">
            Projected Improvement: {projectedImprovement.toFixed(1)}%
          </span>
        </div>

        <div className="text-slate-500">
          Auto-refresh: {autoRefreshOn ? "ON" : "OFF"}
        </div>
      </div>
    </div>
  );
}