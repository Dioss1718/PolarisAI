import { useMemo } from "react";
import { Background, Controls, MiniMap, ReactFlow, MarkerType } from "@xyflow/react";
import "@xyflow/react/dist/style.css";

function riskColor(risk) {
  if (risk >= 8.5) return "#f43f5e";
  if (risk >= 7) return "#f97316";
  if (risk >= 5) return "#facc15";
  return "#22c55e";
}

function cloudBadge(cloud) {
  if (cloud === "AWS") return "◆";
  if (cloud === "AZURE") return "■";
  if (cloud === "GCP") return "▲";
  return "●";
}

function buildPathEdgeSet(path) {
  const set = new Set();
  if (!Array.isArray(path) || path.length < 2) return set;
  for (let i = 0; i < path.length - 1; i++) {
    set.add(`${path[i]}->${path[i + 1]}`);
  }
  return set;
}

function nodeTypeY(type) {
  const map = {
    INTERNET: 60,
    LOAD_BALANCER: 120,
    COMPUTE: 220,
    IAM_ROLE: 340,
    DATABASE: 470,
    STORAGE: 560,
    KUBERNETES: 660,
  };
  return map[type] ?? 260;
}

function cloudBandX(cloud) {
  const map = { AWS: 120, AZURE: 500, GCP: 880 };
  return map[cloud] ?? 120;
}

export default function GraphCanvas({
  nodes = [],
  edges = [],
  nodeIntel = [],
  selectedAttackPath = null,
  onSelectNode,
}) {
  const intelMap = useMemo(
    () => new Map((nodeIntel || []).map((i) => [i.nodeId, i])),
    [nodeIntel]
  );

  const pathEdgeSet = useMemo(() => buildPathEdgeSet(selectedAttackPath), [selectedAttackPath]);
  const pathNodeSet = useMemo(() => new Set(selectedAttackPath || []), [selectedAttackPath]);

  const flowNodes = useMemo(() => {
    const cloudOffsets = {};
    return nodes.map((node) => {
      const intel = intelMap.get(node.id);
      const highlighted = pathNodeSet.has(node.id);
      const xBase = cloudBandX(node.cloud);
      const yBase = nodeTypeY(node.type);
      const seq = cloudOffsets[node.cloud] ?? 0;
      cloudOffsets[node.cloud] = seq + 1;

      const x = xBase + (seq % 2) * 220;
      const y = yBase + Math.floor(seq / 2) * 120;

      const halo = highlighted
        ? "0 0 50px rgba(244,63,94,0.38)"
        : `0 0 ${18 + (intel?.blastRadius || 0) * 1.4}px ${riskColor(node.risk)}25`;

      return {
        id: node.id,
        position: { x, y },
        data: {
          label: `${cloudBadge(node.cloud)} ${node.label}`,
        },
        style: {
          background: highlighted ? "#1e1b4b" : "#08111f",
          color: "white",
          border: `2px solid ${highlighted ? "#f43f5e" : riskColor(node.risk)}`,
          borderRadius: 18,
          width: 220,
          padding: 12,
          fontSize: 12,
          fontWeight: 600,
          boxShadow: halo,
        },
      };
    });
  }, [nodes, intelMap, pathNodeSet]);

  const flowEdges = useMemo(() => {
    return edges.map((edge, index) => {
      const highlighted = pathEdgeSet.has(`${edge.from}->${edge.to}`);
      return {
        id: `e-${index}`,
        source: edge.from,
        target: edge.to,
        label: edge.type,
        animated: highlighted,
        type: "smoothstep",
        markerEnd: {
          type: MarkerType.ArrowClosed,
          color: highlighted ? "#f43f5e" : "#64748b",
        },
        style: {
          stroke: highlighted ? "#f43f5e" : "#64748b",
          strokeWidth: highlighted ? 3.5 : 1.6,
          opacity: highlighted ? 1 : 0.75,
        },
        labelStyle: {
          fill: highlighted ? "#fecdd3" : "#cbd5e1",
          fontSize: 10,
        },
      };
    });
  }, [edges, pathEdgeSet]);

  return (
    <div className="h-full min-h-[420px] w-full overflow-hidden rounded-xl bg-slate-950">
      <ReactFlow
        nodes={flowNodes}
        edges={flowEdges}
        fitView
        style={{ width: "100%", height: "100%" }}
        onNodeClick={(_, node) => onSelectNode?.(nodes.find((n) => n.id === node.id))}
      >
        <MiniMap pannable zoomable />
        <Controls showInteractive={false} />
        <Background color="#1e293b" gap={18} size={1} />
      </ReactFlow>
    </div>
  );
}