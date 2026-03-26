import { useMemo, useState } from "react";
import TopBar from "../components/TopBar";
import SimulationStudio from "../components/SimulationStudio";
import GraphCanvas from "../components/GraphCanvas";
import NodeDrawer from "../components/NodeDrawer";
import DecisionRail from "../components/DecisionRail";
import RecommendationPanel from "../components/RecommendationPanel";
import ExplainabilityPanel from "../components/ExplainabilityPanel";
import ForecastPanel from "../components/ForecastPanel";
import GitOpsPanel from "../components/GitOpsPanel";
import FeedbackPanel from "../components/FeedbackPanel";
import CopilotPanel from "../components/CopilotPanel";
import { runPipeline } from "../api/client";

export default function ConsolePage() {
  const [scenario, setScenario] = useState("FULL_CHAOS");
  const [seed, setSeed] = useState(42);
  const [loading, setLoading] = useState(false);
  const [state, setState] = useState({
    nodes: [],
    edges: [],
    recommendations: [],
    explanations: [],
    forecasts: [],
    feedback: {},
    gitops: { status: "skipped", message: "No run yet", prs: [] },
    attackPaths: [],
    stages: [],
    scenario: "",
    seed: 42,
  });
  const [selectedNodeId, setSelectedNodeId] = useState(null);

  const selectedNode = useMemo(
    () => state.nodes.find((n) => n.id === selectedNodeId) || null,
    [state.nodes, selectedNodeId]
  );

  const selectedRecommendation = useMemo(
    () => state.recommendations.find((r) => r.nodeId === selectedNodeId) || null,
    [state.recommendations, selectedNodeId]
  );

  const selectedExplanation = useMemo(
    () => state.explanations.find((e) => e.nodeId === selectedNodeId) || null,
    [state.explanations, selectedNodeId]
  );

  const selectedForecast = useMemo(
    () => state.forecasts.find((f) => f.nodeId === selectedNodeId) || null,
    [state.forecasts, selectedNodeId]
  );

  const doRun = async (payload) => {
    setLoading(true);
    try {
      const data = await runPipeline(payload);
      setState(data);
      if (data.nodes?.length) {
        setSelectedNodeId(data.nodes[0].id);
      }
    } catch (err) {
      const msg = err?.response?.data?.error || err.message || "Pipeline run failed";
      alert(msg);
    } finally {
      setLoading(false);
    }
  };

  const runScenario = () => doRun({ scenario, seed });
  const runManual = (manualData) => doRun({ scenario: "MANUAL", seed, manualData });

  const selectByRecommendation = (rec) => {
    setSelectedNodeId(rec.nodeId);
  };

  const selectNodeById = (nodeId) => {
    setSelectedNodeId(nodeId);
  };

  return (
    <div className="min-h-screen px-4 py-5 lg:px-6">
      <div className="mx-auto max-w-[1800px] space-y-5">
        <TopBar
          onRun={runScenario}
          loading={loading}
          scenario={scenario}
          setScenario={setScenario}
          seed={seed}
          setSeed={setSeed}
        />

        <div className="grid grid-cols-1 gap-5 2xl:grid-cols-[1.15fr_0.85fr]">
          <GraphCanvas
            nodes={state.nodes}
            edges={state.edges}
            onSelectNode={(node) => setSelectedNodeId(node?.id)}
          />

          <NodeDrawer
            node={selectedNode}
            recommendation={selectedRecommendation}
            explanation={selectedExplanation}
            forecast={selectedForecast}
          />
        </div>

        <div className="grid grid-cols-1 gap-5 xl:grid-cols-[0.95fr_1.05fr]">
          <SimulationStudio onRunManual={runManual} loading={loading} />
          <DecisionRail stages={state.stages} />
        </div>

        <div className="grid grid-cols-1 gap-5 xl:grid-cols-[1fr_1fr]">
          <RecommendationPanel
            recommendations={state.recommendations}
            onSelect={selectByRecommendation}
          />
          <ExplainabilityPanel
            explanations={state.explanations}
            onSelect={selectNodeById}
          />
        </div>

        <ForecastPanel forecasts={state.forecasts} onSelect={selectNodeById} />

        <div className="grid grid-cols-1 gap-5 xl:grid-cols-[1fr_1fr]">
          <GitOpsPanel gitops={state.gitops} />
          <FeedbackPanel feedback={state.feedback} />
        </div>

        <CopilotPanel state={state} onSelectNode={selectNodeById} />
      </div>
    </div>
  );
}