import React, { useEffect, useMemo, useState } from "react";
import { Background, Controls, MiniMap, ReactFlow } from "@xyflow/react";
import "@xyflow/react/dist/style.css";

const riskColor = (risk) => {
  if (risk >= 8) return "#ef4444";
  if (risk >= 6) return "#f97316";
  if (risk >= 4) return "#eab308";
  return "#22c55e";
};

export default function GraphCanvas({ nodes = [], edges = [], onSelectNode }) {
  const [rfKey, setRfKey] = useState(0);

  useEffect(() => {
    setRfKey((k) => k + 1);
  }, [nodes, edges]);

  const flowNodes = useMemo(() => {
    return nodes.map((node, index) => ({
      id: node.id,
      position: {
        x: 120 + (index % 4) * 240,
        y: 100 + Math.floor(index / 4) * 180,
      },
      data: {
        label: `${node.label} · ${node.cloud}`,
      },
      style: {
        background: "#0f172a",
        color: "#fff",
        border: `2px solid ${riskColor(node.risk)}`,
        borderRadius: 16,
        padding: 12,
        width: 180,
        boxShadow: `0 0 24px ${riskColor(node.risk)}33`,
        fontSize: 13,
        fontWeight: 500,
        textAlign: "center",
      },
    }));
  }, [nodes]);

  const flowEdges = useMemo(() => {
    return edges.map((edge, i) => ({
      id: `edge-${i}`,
      source: edge.from,
      target: edge.to,
      label: edge.type,
      style: { stroke: "#64748b", strokeWidth: 1.5 },
      labelStyle: { fill: "#cbd5e1", fontSize: 10, fontWeight: 500 },
      labelBgStyle: {
        fill: "#e5e7eb",
        fillOpacity: 0.95,
        rx: 6,
        ry: 6,
      },
    }));
  }, [edges]);

  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-3 h-[560px]">
      <div className="mb-3">
        <h2 className="text-lg font-semibold">Unified Cloud Graph</h2>
        <p className="text-sm text-slate-400">
          Risk-colored nodes, governance decisions, and attack-path aware topology.
        </p>
      </div>

      <div className="h-[490px] rounded-xl overflow-hidden bg-slate-950">
        {flowNodes.length === 0 ? (
          <div className="flex h-full items-center justify-center text-sm text-slate-500">
            Run governance to load graph topology.
          </div>
        ) : (
          <ReactFlow
            key={rfKey}
            nodes={flowNodes}
            edges={flowEdges}
            fitView
            onNodeClick={(_, n) => {
              const selected = nodes.find((x) => x.id === n.id);
              if (onSelectNode) onSelectNode(selected);
            }}
          >
            <MiniMap
              pannable
              zoomable
              position="bottom-right"
              nodeColor={(node) => {
                const original = nodes.find((n) => n.id === node.id);
                return original ? riskColor(original.risk) : "#64748b";
              }}
              maskColor="rgba(2, 6, 23, 0.72)"
              style={{
                backgroundColor: "#0f172a",
                border: "1px solid #334155",
                borderRadius: 14,
                boxShadow: "0 10px 30px rgba(0,0,0,0.35)",
              }}
            />

            <Controls
              position="bottom-left"
              showInteractive={false}
              style={{
                background: "#0f172a",
                border: "1px solid #334155",
                borderRadius: "14px",
                boxShadow: "0 10px 30px rgba(0,0,0,0.35)",
                overflow: "hidden",
              }}
            />

            <Background
              color="#1e293b"
              gap={18}
              size={1}
            />
          </ReactFlow>
        )}
      </div>
    </div>
  );
}