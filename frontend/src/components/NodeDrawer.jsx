export default function NodeDrawer({ node, recommendation, forecast }) {
  if (!node) {
    return (
      <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4 min-h-[560px]">
        <h2 className="text-lg font-semibold">Node Detail</h2>
        <p className="mt-4 text-sm text-slate-400">
          Select a node from the graph or a recommendation from the governance feed.
        </p>
      </div>
    );
  }

  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4 min-h-[560px] overflow-auto">
      <h2 className="text-lg font-semibold">{node.label}</h2>
      <div className="mt-3 grid grid-cols-2 gap-3 text-sm">
        <Info label="Node ID" value={node.id} />
        <Info label="Cloud" value={node.cloud} />
        <Info label="Type" value={node.type} />
        <Info label="Region" value={node.region} />
        <Info label="Environment" value={node.environment} />
        <Info label="Exposure" value={node.exposure} />
        <Info label="Risk" value={node.risk} />
        <Info label="Cost" value={node.cost} />
        <Info label="Utilization" value={node.utilization} />
        <Info label="Final Action" value={node.finalAction || "-"} />
        <Info label="Status" value={node.status || "-"} />
      </div>

      {recommendation && (
        <section className="mt-6">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-sky-300">Governance Decision</h3>
          <div className="mt-2 rounded-xl bg-slate-950/70 p-3 text-sm">
            <div><span className="text-slate-400">Action:</span> {recommendation.finalAction}</div>
            <div><span className="text-slate-400">Status:</span> {recommendation.status}</div>
            <div><span className="text-slate-400">Score:</span> {recommendation.score}</div>
            <div className="mt-2 text-slate-300">{recommendation.reason}</div>
          </div>
        </section>
      )}

      {forecast && (
        <section className="mt-6">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-emerald-300">Forecast</h3>
          <div className="mt-2 rounded-xl bg-slate-950/70 p-3 text-sm">
            <div>Current: {forecast.currentCost}</div>
            <div>30d: {forecast.forecast30}</div>
            <div>90d: {forecast.forecast90}</div>
            <div>Shock: {forecast.billShock ? "YES" : "NO"}</div>
            {forecast.shockReason ? <div className="mt-2 text-slate-400">{forecast.shockReason}</div> : null}
          </div>
        </section>
      )}

      
    </div>
  );
}

function Info({ label, value }) {
  return (
    <div className="rounded-xl bg-slate-950/60 p-3">
      <div className="text-xs uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 text-slate-100">{String(value)}</div>
    </div>
  );
}