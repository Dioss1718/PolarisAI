import GraphCanvas from "./GraphCanvas";
import GraphLegend from "./GraphLegend";

export default function GraphPanel({
  nodes,
  edges,
  attackPathCount,
  selectedAttackPath,
  nodeIntel,
  onSelectNode,
}) {
  return (
    <div className="h-full min-h-0 rounded-2xl border border-white/10 bg-white/[0.04] p-3">
      <div className="mb-2 flex flex-wrap items-start justify-between gap-3">
        <div>
          <div className="text-lg font-semibold">Unified Cloud Graph</div>
          <div className="text-sm text-slate-400">
            Attack-path highlighting, severity halos, cloud grouping, and propagation-aware topology.
          </div>
        </div>

        <div className="flex items-center gap-2">
          <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-xs text-slate-300">
            Attack Paths: {attackPathCount}
          </div>
          <div className="rounded-full border border-white/10 bg-slate-950/60 px-3 py-1 text-xs text-slate-300">
            Graph Hero View
          </div>
        </div>
      </div>

      <div className="mb-2">
        <GraphLegend />
      </div>

      <div className="h-[calc(100%-72px)] min-h-[460px] w-full">
        <GraphCanvas
          nodes={nodes}
          edges={edges}
          nodeIntel={nodeIntel}
          selectedAttackPath={selectedAttackPath}
          onSelectNode={onSelectNode}
        />
      </div>
    </div>
  );
}