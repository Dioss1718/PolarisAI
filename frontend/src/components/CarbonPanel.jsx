export default function CarbonPanel({ summary, projectedSummary }) {
  const before = Number(summary?.currentCarbonTotal ?? 0);
  const after = Number(projectedSummary?.projectedCarbonTotal ?? 0);
  const reduction = Number(projectedSummary?.carbonReductionPct ?? 0);
  const greenScore = Number(projectedSummary?.greenScore ?? 0);

  const topSources =
    summary?.topCarbonSources ||
    projectedSummary?.topCarbonSources ||
    [];

  const actionImpact =
    projectedSummary?.carbonActionImpact ||
    [];

  return (
    <div className="space-y-4">
      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="text-lg font-semibold">Carbon Summary</div>
        <div className="mt-4 grid gap-3 md:grid-cols-4">
          <Metric label="Before" value={before.toFixed(1)} />
          <Metric label="After" value={after.toFixed(1)} />
          <Metric label="Reduction %" value={`${reduction.toFixed(1)}%`} />
          <Metric label="Green Score" value={greenScore.toFixed(1)} />
        </div>
      </div>

      {Array.isArray(topSources) && topSources.length > 0 ? (
        <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
          <div className="text-lg font-semibold">Top Carbon Sources</div>
          <div className="mt-4 grid gap-3">
            {topSources.slice(0, 3).map((item, index) => (
              <div
                key={`${item.nodeId || item.name || index}`}
                className="rounded-xl border border-white/10 bg-slate-900/70 p-3"
              >
                <div className="font-medium">
                  {item.nodeId || item.name || `Source ${index + 1}`}
                </div>
                <div className="mt-1 text-sm text-slate-400">
                  Contribution: {Number(item.percentContribution || item.percentage || 0).toFixed(1)}%
                </div>
              </div>
            ))}
          </div>
        </div>
      ) : null}

      {Array.isArray(actionImpact) && actionImpact.length > 0 ? (
        <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
          <div className="text-lg font-semibold">Action Impact</div>
          <div className="mt-4 grid gap-3">
            {actionImpact.map((item, index) => (
              <div
                key={`${item.action || index}`}
                className="rounded-xl border border-white/10 bg-slate-900/70 p-3"
              >
                <div className="font-medium">{item.action || `Action ${index + 1}`}</div>
                <div className="mt-1 text-sm text-slate-400">
                  Carbon Reduction: {Number(item.carbonReduction || item.reduction || 0).toFixed(1)}
                </div>
              </div>
            ))}
          </div>
        </div>
      ) : null}
    </div>
  );
}

function Metric({ label, value }) {
  return (
    <div className="rounded-xl bg-slate-900/70 p-3">
      <div className="text-[11px] uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 text-lg font-semibold">{value}</div>
    </div>
  );
}