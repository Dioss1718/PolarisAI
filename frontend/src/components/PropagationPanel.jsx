export default function PropagationPanel({ nodeIntel = [], edgeInfluence = [] }) {
  const topNodes = [...nodeIntel]
    .sort((a, b) => Number(b.riskInfluence || 0) - Number(a.riskInfluence || 0))
    .slice(0, 8);

  const topEdges = [...edgeInfluence].slice(0, 10);

  return (
    <div className="grid gap-4 md:grid-cols-[380px_1fr]">
      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="text-lg font-semibold">Risk Propagation Hotspots</div>
        <div className="mt-3 space-y-3">
          {topNodes.map((n) => (
            <div key={n.nodeId} className="rounded-xl border border-white/10 bg-slate-900/70 p-3">
              <div className="flex items-center justify-between gap-3">
                <div className="font-medium">{n.nodeId}</div>
                <div className="text-sm text-rose-300">
                  Influence {Number(n.riskInfluence || 0).toFixed(2)}
                </div>
              </div>
              <div className="mt-1 text-xs text-slate-500">
                Blast Radius {n.blastRadius} · Attack Paths {n.attackPathCount}
              </div>
              <div className="mt-2 text-sm text-slate-300">{n.why}</div>
            </div>
          ))}
        </div>
      </div>

      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="text-lg font-semibold">Propagation Edges</div>
        <div className="mt-4 grid gap-3">
          {topEdges.map((e, idx) => (
            <div key={`${e.from}-${e.to}-${idx}`} className="rounded-xl border border-white/10 bg-slate-900/70 p-3">
              <div className="flex items-center justify-between gap-3">
                <div className="font-medium">
                  {e.from} <span className="text-slate-500">→</span> {e.to}
                </div>
                <div className="text-sm text-amber-300">
                  {Number(e.influence || 0).toFixed(2)}
                </div>
              </div>
              <div className="mt-2 text-sm text-slate-300">{e.reason}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}