import React, { useEffect, useMemo, useState } from "react";
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
import ServiceStatusBar from "../components/ServiceStatusBar";
import { getServiceHealth, runPipeline } from "../api/client";

export default function ConsolePage() {
  const [scenario, setScenario] = useState("FULL_CHAOS");
  const [seed, setSeed] = useState(42);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [services, setServices] = useState([]);
  const [state, setState] = useState({
    nodes: [],
    edges: [],
    recommendations: [],
    explanations: [],
    forecasts: [],
    feedback: {},
    gitops: { status: "idle", message: "No run yet", prs: [] },
    attackPaths: [],
    stages: [],
    scenario: "",
    seed: 42,
    generatedAt: "",
  });
  const [selectedNodeId, setSelectedNodeId] = useState(null);

  useEffect(() => {
    refreshServices();
  }, []);

  const refreshServices = async () => {
    try {
      const data = await getServiceHealth();
      setServices(Array.isArray(data)? data : data.services || []);
    } catch {
      setServices([
        { name: "Governance API", status: "down" },
        { name: "AI Engine", status: "down" },
        { name: "Forecast Engine", status: "down" },
      ]);
    }
  };

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
    setErrorMessage("");

    try {
      const data = await runPipeline(payload);
      setState({
  ...data,
  nodes: Array.isArray(data.nodes) ? data.nodes : [],
  edges: Array.isArray(data.edges) ? data.edges : [],
  recommendations: Array.isArray(data.recommendations) ? data.recommendations : [],
  explanations: Array.isArray(data.explanations) ? data.explanations : [],
  forecasts: Array.isArray(data.forecasts) ? data.forecasts : [],
});
      if (data.nodes?.length) {
        setSelectedNodeId(data.nodes[0].id);
      }
      await refreshServices();
    } catch (err) {
      const msg = err?.response?.data?.error || err.message || "Pipeline run failed";
      setErrorMessage(msg);
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
      <div className="mx-auto max-w-[1840px] space-y-5">
        <TopBar
          onRun={runScenario}
          loading={loading}
          scenario={scenario}
          setScenario={setScenario}
          seed={seed}
          setSeed={setSeed}
        />

        <ServiceStatusBar
          services={services}
          generatedAt={state.generatedAt}
          scenario={state.scenario}
          seed={state.seed}
        />

        {errorMessage ? (
          <div className="rounded-2xl border border-rose-500/25 bg-rose-500/10 p-4 text-sm text-rose-200 shadow-glow">
            <div className="font-semibold">Pipeline Error</div>
            <div className="mt-1">{errorMessage}</div>
          </div>
        ) : null}

        {loading ? (
          <div className="rounded-2xl border border-sky-500/20 bg-sky-500/10 p-4 text-sm text-sky-200 shadow-glow">
            Governance execution is running. Large AI or GitOps stages may take longer depending on environment state.
          </div>
        ) : null}

        <div className="grid grid-cols-1 gap-5 2xl:grid-cols-[1.15fr_0.85fr]">
          <GraphCanvas
            nodes={state.nodes}
            edges={state.edges}
            attackPathCount={state.attackPaths?.length || 0}
            onSelectNode={(node) => setSelectedNodeId(node?.id)}
          />

          <NodeDrawer
            node={selectedNode}
            recommendation={selectedRecommendation}
            
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