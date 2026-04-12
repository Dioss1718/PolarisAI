import { useEffect, useMemo, useState } from "react";
import {
  clearSession,
  compareWhatIf,
  getMe,
  getServiceHealth,
  getSession,
  getState,
  runPipeline,
} from "../api/client";
import { deriveMetrics } from "../api/selectors";
import AppShell from "../layout/AppShell";
import HeaderBar from "../layout/HeaderBar";
import MetaBar from "../layout/MetaBar";
import LoginDialog from "../components/LoginDialog";
import CopilotDrawer from "../components/CopilotDrawer";
import CopilotLauncher from "../components/CopilotLauncher";
import BottomStatusBar from "../components/BottomStatusBar";
import AnalysisWorkspacePage from "./AnalysisWorkspacePage";
import SimulationStudio from "../components/SimulationStudio";

const initialState = {
  nodes: [],
  edges: [],
  recommendations: [],
  explanations: [],
  forecasts: [],
  feedback: {},
  gitops: { status: "idle", message: "No run yet", prs: [] },
  attackPaths: [],
  attackMetrics: null,
  stages: [],
  scenario: "",
  seed: 42,
  generatedAt: "",
  summary: null,
  projectedSummary: null,
  alerts: [],
  nodeIntel: [],
  edgeInfluence: [],
  negotiations: [],
};

function isUnauthorized(err) {
  return err?.response?.status === 401;
}

export default function ConsolePage() {
  const [session, setSession] = useState(getSession());
  const [scenario, setScenario] = useState("FULL_CHAOS");
  const [seed, setSeed] = useState(42);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [services, setServices] = useState([]);
  const [state, setState] = useState(initialState);
  const [selectedNodeId, setSelectedNodeId] = useState(null);
  const [selectedPathIndex, setSelectedPathIndex] = useState(null);
  const [loginOpen, setLoginOpen] = useState(false);
  const [copilotOpen, setCopilotOpen] = useState(false);
  const [workspaceOpen, setWorkspaceOpen] = useState(false);
  const [workspaceTab, setWorkspaceTab] = useState("governance");
  const [polling, setPolling] = useState(false);
  const [hasRun, setHasRun] = useState(false);
  const [pendingRunAfterLogin, setPendingRunAfterLogin] = useState(false);
  const [highlightGraphNav, setHighlightGraphNav] = useState(false);

  const invalidateSession = (
    message = "Session expired or unauthorized. Please log in again."
  ) => {
    clearSession();
    setSession(null);
    setPolling(false);
    setHasRun(false);
    setServices([]);
    setState(initialState);
    setSelectedNodeId(null);
    setSelectedPathIndex(null);
    setWorkspaceOpen(false);
    setPendingRunAfterLogin(false);
    setHighlightGraphNav(false);
    setErrorMessage(message);
    setLoginOpen(true);
  };

  useEffect(() => {
    let stopped = false;

    const verifyStartupSession = async () => {
      const existing = getSession();

      if (!existing) {
        if (!stopped) {
          setSession(null);
          setLoginOpen(true);
        }
        return;
      }

      try {
        const me = await getMe();
        if (!stopped) {
          setSession(me || existing);
          setLoginOpen(false);
        }
      } catch (err) {
        if (isUnauthorized(err)) {
          if (!stopped) {
            invalidateSession("Please log in to continue.");
          }
        } else if (!stopped) {
          setSession(existing);
        }
      }
    };

    verifyStartupSession();

    return () => {
      stopped = true;
    };
  }, []);

  useEffect(() => {
    if (!session) {
      setPolling(false);
      setHasRun(false);
      setLoginOpen(true);
      return;
    }

    let stopped = false;

    const loadHealthOnly = async () => {
      try {
        const serviceData = await getServiceHealth();
        if (!stopped) {
          setServices(Array.isArray(serviceData?.services) ? serviceData.services : []);
        }
      } catch {}
    };

    loadHealthOnly();

    return () => {
      stopped = true;
    };
  }, [session]);

  useEffect(() => {
    if (!session || !hasRun || !polling || loading) return;

    let stopped = false;

    const loadLatestState = async () => {
      try {
        const data = await getState({ scenario, seed });
        if (!stopped && data) {
          hydrateState(data);
        }
      } catch (err) {
        if (isUnauthorized(err) && !stopped) {
          invalidateSession();
        }
      }
    };

    const timer = setInterval(loadLatestState, 5000);

    return () => {
      stopped = true;
      clearInterval(timer);
    };
  }, [session, hasRun, polling, loading, scenario, seed]);

  const selectedNode = useMemo(
    () => state.nodes.find((n) => n.id === selectedNodeId) || null,
    [state.nodes, selectedNodeId]
  );

  const selectedRecommendation = useMemo(
    () => state.recommendations.find((r) => r.nodeId === selectedNodeId) || null,
    [state.recommendations, selectedNodeId]
  );

  const selectedForecast = useMemo(
    () => state.forecasts.find((f) => f.nodeId === selectedNodeId) || null,
    [state.forecasts, selectedNodeId]
  );

  const selectedAttackPath = useMemo(() => {
    if (selectedPathIndex === null || selectedPathIndex === undefined) return null;
    return state.attackPaths?.[selectedPathIndex] || null;
  }, [state.attackPaths, selectedPathIndex]);

  const metrics = useMemo(() => deriveMetrics(state), [state]);

  const hydrateState = (data) => {
    setState((prev) => ({
      ...prev,
      ...data,
      nodes: Array.isArray(data?.nodes) ? data.nodes : [],
      edges: Array.isArray(data?.edges) ? data.edges : [],
      recommendations: Array.isArray(data?.recommendations) ? data.recommendations : [],
      explanations: Array.isArray(data?.explanations) ? data.explanations : [],
      forecasts: Array.isArray(data?.forecasts) ? data.forecasts : [],
      attackPaths: Array.isArray(data?.attackPaths) ? data.attackPaths : [],
      attackMetrics: data?.attackMetrics || null,
      summary: data?.summary || null,
      projectedSummary: data?.projectedSummary || null,
      stages: Array.isArray(data?.stages) ? data.stages : [],
      alerts: Array.isArray(data?.alerts) ? data.alerts : [],
      nodeIntel: Array.isArray(data?.nodeIntel) ? data.nodeIntel : [],
      edgeInfluence: Array.isArray(data?.edgeInfluence) ? data.edgeInfluence : [],
      negotiations: Array.isArray(data?.negotiations) ? data.negotiations : [],
      gitops: data?.gitops || { status: "idle", message: "No run yet", prs: [] },
    }));

    if (Array.isArray(data?.nodes) && data.nodes.length) {
      setSelectedNodeId((prev) => prev || data.nodes[0].id);
    }

    setHasRun(true);
  };

  const executeRun = async (manualData = null) => {
    setLoading(true);
    setPolling(false);
    setErrorMessage("");

    try {
      const payload = manualData
        ? { scenario: "MANUAL", seed, manualData }
        : { scenario, seed };

      const data = await runPipeline(payload);
      hydrateState(data);
      setPolling(true);
      setHighlightGraphNav(true);
    } catch (err) {
      if (isUnauthorized(err)) {
        invalidateSession();
      } else {
        setErrorMessage(
          err?.response?.data?.error ||
            err.message ||
            (manualData ? "Manual simulation failed" : "Pipeline run failed")
        );
      }
    } finally {
      setLoading(false);
    }
  };

  const runScenario = async () => {
    if (!session) {
      setPendingRunAfterLogin(true);
      setLoginOpen(true);
      return;
    }
    await executeRun(null);
  };

  const runManual = async (manualData) => {
    if (!session) {
      setPendingRunAfterLogin(true);
      setLoginOpen(true);
      return;
    }
    await executeRun(manualData);
  };

  const runWhatIfCompare = async (payload) => {
    if (!session) {
      setLoginOpen(true);
      throw new Error("Please log in to run what-if comparison.");
    }

    return await compareWhatIf(payload);
  };

  const handleLoggedIn = async (newSession) => {
    setErrorMessage("");
    setLoginOpen(false);

    try {
      const me = await getMe();
      const finalSession = me || newSession;
      setSession(finalSession);

      if (pendingRunAfterLogin) {
        setPendingRunAfterLogin(false);
        setTimeout(() => {
          executeRun(null);
        }, 0);
      }
    } catch (err) {
      if (isUnauthorized(err)) {
        invalidateSession("Login succeeded locally, but backend rejected the session.");
      } else {
        setSession(newSession);
        if (pendingRunAfterLogin) {
          setPendingRunAfterLogin(false);
          setTimeout(() => {
            executeRun(null);
          }, 0);
        }
      }
    }
  };

  const logout = () => {
    clearSession();
    setSession(null);
    setPolling(false);
    setHasRun(false);
    setLoginOpen(true);
    setHighlightGraphNav(false);
  };

  const openWorkspaceTab = (tab) => {
    if (tab === "graph") {
      setHighlightGraphNav(false);
    }
    setWorkspaceTab(tab);
    setWorkspaceOpen(true);
  };

  const simulationAllowed = session?.features?.SIMULATION_STUDIO === "FULL";

  return (
    <AppShell>
      {workspaceOpen ? (
        <AnalysisWorkspacePage
          activeTab={workspaceTab}
          setActiveTab={setWorkspaceTab}
          state={state}
          session={session}
          onBack={() => setWorkspaceOpen(false)}
          onSelectNode={setSelectedNodeId}
          onSelectPath={setSelectedPathIndex}
          selectedPathIndex={selectedPathIndex}
          selectedNode={selectedNode}
          selectedRecommendation={selectedRecommendation}
          selectedForecast={selectedForecast}
          selectedAttackPath={selectedAttackPath}
          onRunManual={runManual}
          loading={loading}
          simulationAllowed={simulationAllowed}
          onRefreshState={async () => {
            const data = await getState({
              scenario: state.scenario || scenario,
              seed: state.seed ?? seed,
            });
            if (data) {
              hydrateState(data);
            }
          }}
        />
      ) : (
        <div className="pb-24">
          <div className="space-y-3">
            <HeaderBar
              scenario={scenario}
              setScenario={setScenario}
              seed={seed}
              setSeed={setSeed}
              onRun={runScenario}
              loading={loading}
              onOpenLogin={() => setLoginOpen(true)}
              session={session}
              onLogout={logout}
              onOpenWorkspace={openWorkspaceTab}
            />

            <MetaBar
              metrics={metrics}
              scenario={state.scenario || scenario}
              seed={state.seed ?? seed}
              services={services}
              gitops={state.gitops}
              onOpenWorkspace={openWorkspaceTab}
              highlightGraphNav={highlightGraphNav}
            />

            <div className="grid items-start gap-3 xl:grid-cols-[minmax(0,0.52fr)_minmax(700px,0.48fr)]">
              <TopGovernanceActionsPanel
                recommendations={state.recommendations}
                onOpenWorkspace={openWorkspaceTab}
              />

              <SimulationStudio
              state={state}
              onCompare={runWhatIfCompare}
              loading={loading}
              allowed={simulationAllowed}
              scenario={state.scenario || scenario}
              seed={state.seed ?? seed}
            />
            </div>
          </div>
        </div>
      )}

      <BottomStatusBar
        session={session}
        services={services}
        metrics={metrics}
        autoRefreshOn={polling}
      />

      <CopilotLauncher onClick={() => setCopilotOpen(true)} />
      <CopilotDrawer
        open={copilotOpen}
        onClose={() => setCopilotOpen(false)}
        state={state}
        onSelectNode={setSelectedNodeId}
      />
      <LoginDialog
        open={loginOpen}
        onClose={() => {
          if (session) setLoginOpen(false);
        }}
        onLoggedIn={handleLoggedIn}
      />

      {errorMessage ? (
        <div className="fixed bottom-16 left-4 z-[120] max-w-lg rounded-xl border border-rose-500/30 bg-rose-500/10 px-4 py-3 text-sm text-rose-200 shadow-lg backdrop-blur">
          {errorMessage}
        </div>
      ) : null}
    </AppShell>
  );
}

function TopGovernanceActionsPanel({ recommendations = [], onOpenWorkspace }) {
  const top = [...(recommendations || [])]
    .sort((a, b) => Number(b.risk || 0) - Number(a.risk || 0))
    .slice(0, 6);

  return (
    <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="flex items-center justify-between gap-3">
        <div>
          <div className="text-lg font-semibold">Top Governance Actions</div>
          <div className="text-sm text-slate-400">
            Highest-priority governance decisions from the latest run.
          </div>
        </div>

        <button
          onClick={() => onOpenWorkspace?.("governance")}
          className="rounded-xl border border-sky-500/20 bg-sky-500/10 px-4 py-2 text-sm text-sky-200"
        >
          Open Governance Workspace
        </button>
      </div>

      <div className="mt-4 h-[520px] overflow-auto pr-1">
        {top.length === 0 ? (
          <div className="rounded-xl border border-white/10 bg-slate-950/60 p-4 text-sm text-slate-400">
            No governance actions available yet.
          </div>
        ) : (
          <div className="grid gap-3">
            {top.map((r) => (
              <div
                key={`${r.nodeId}-${r.finalAction}`}
                className="rounded-xl border border-white/10 bg-slate-950/60 p-4"
              >
                <div className="flex items-start justify-between gap-3">
                  <div>
                    <div className="font-semibold text-slate-100">{r.nodeId}</div>
                    <div className="mt-1 text-sm text-slate-400">
                      {r.finalAction} · {r.status}
                    </div>
                  </div>

                  <div className="rounded-full border border-white/10 bg-slate-900/70 px-3 py-1 text-xs text-slate-300">
                    Risk {Number(r.risk || 0).toFixed(2)}
                  </div>
                </div>

                <div className="mt-3 whitespace-pre-wrap break-words text-sm text-slate-300">
                  {r.reason}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}