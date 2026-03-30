function info(label, value) {
  return (
    <div className="rounded-xl bg-slate-950/60 p-3">
      <div className="text-xs uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 break-words">{String(value ?? "—")}</div>
    </div>
  );
}

export default function NodeDrawer({
  node,
  recommendation,
  forecast,
  nodeIntel = [],
}) {
  if (!node) {
    return (
      <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        Select a node to inspect its governance reasoning.
      </div>
    );
  }

  const intel = (nodeIntel || []).find((n) => n.nodeId === node.id);

  return (
    <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="flex items-start justify-between gap-3">
        <div>
          <div className="text-lg font-semibold">{node.label}</div>
          <div className="mt-1 text-xs text-slate-500">{node.id}</div>
        </div>
        <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-xs">
          Risk {Number(node.risk || 0).toFixed(2)}
        </div>
      </div>

      <div className="mt-4 grid grid-cols-2 gap-3 text-sm">
        {info("Cloud", node.cloud)}
        {info("Type", node.type)}
        {info("Region", node.region)}
        {info("Environment", node.environment)}
        {info("Exposure", node.exposure)}
        {info("Cost", node.cost)}
        {info("Utilization", node.utilization)}
        {info("Final Action", node.finalAction || "-")}
        {info("Status", node.status || "-")}
        {info("Blast Radius", intel?.blastRadius)}
        {info("Risk Influence", intel?.riskInfluence)}
        {info("Attack Path Count", intel?.attackPathCount)}
      </div>

      {intel?.why ? (
        <section className="mt-5 rounded-xl bg-slate-950/60 p-4 text-sm">
          <div className="text-xs uppercase tracking-wider text-sky-300">Why this node matters</div>
          <div className="mt-2 text-slate-300">{intel.why}</div>
        </section>
      ) : null}

      {intel?.affectedNodes?.length ? (
        <section className="mt-5 rounded-xl bg-slate-950/60 p-4 text-sm">
          <div className="text-xs uppercase tracking-wider text-rose-300">Affected downstream assets</div>
          <div className="mt-3 flex flex-wrap gap-2">
            {intel.affectedNodes.map((n) => (
              <span
                key={n}
                className="rounded-full border border-white/10 bg-slate-900/70 px-3 py-1 text-xs text-slate-300"
              >
                {n}
              </span>
            ))}
          </div>
        </section>
      ) : null}

      {recommendation ? (
        <section className="mt-5 rounded-xl bg-slate-950/60 p-4 text-sm">
          <div className="text-xs uppercase tracking-wider text-sky-300">Governance Decision</div>
          <div className="mt-2 grid grid-cols-2 gap-3">
            {info("Selected Action", recommendation.finalAction)}
            {info("Safety", recommendation.safetyLevel)}
            {info("Confidence", recommendation.confidence)}
            {info("Score", recommendation.score)}
          </div>
          <div className="mt-3 text-slate-300">{recommendation.reason}</div>
          <div className="mt-3 grid gap-3">
            {info("GitOps Path", recommendation.gitOpsPath)}
            {info("Rollback Path", recommendation.rollbackPath)}
          </div>
        </section>
      ) : null}

      {forecast ? (
        <section className="mt-5 rounded-xl bg-slate-950/60 p-4 text-sm">
          <div className="text-xs uppercase tracking-wider text-emerald-300">FinOps Outlook</div>
          <div className="mt-2 grid grid-cols-2 gap-3">
            {info("Current", forecast.currentCost)}
            {info("30d", forecast.forecast30)}
            {info("90d", forecast.forecast90)}
            {info("Bill Shock", forecast.billShock ? "YES" : "NO")}
          </div>
          {forecast.shockReason ? <div className="mt-3 text-slate-300">{forecast.shockReason}</div> : null}
        </section>
      ) : null}
    </div>
  );
}