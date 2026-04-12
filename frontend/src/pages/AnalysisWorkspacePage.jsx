import RecommendationList from "../components/RecommendationList";
import ExplainabilityList from "../components/ExplainabilityList";
import ForecastPanel from "../components/ForecastPanel";
import GitOpsPanel from "../components/GitOpsPanel";
import FeedbackPanel from "../components/FeedbackPanel";
import CarbonPanel from "../components/CarbonPanel";
import WorkspaceTabs from "../components/WorkspaceTabs";
import PropagationPanel from "../components/PropagationPanel";
import AttackPathsPanel from "../components/AttackPathsPanel";
import GraphPanel from "../components/GraphPanel";
import GraphActivityPanel from "../components/GraphActivityPanel";
import NodeDrawer from "../components/NodeDrawer";

function NegotiationPanel({ negotiations = [] }) {
  return (
    <div className="space-y-4">
      {negotiations.map((n) => (
        <div key={n.nodeId} className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
          <div className="flex items-center justify-between">
            <div>
              <div className="font-medium">{n.nodeId}</div>
              <div className="text-xs text-slate-500">Selected: {n.selectedAction}</div>
            </div>
            <div className="text-sm text-sky-300">Score {Number(n.selectedScore || 0).toFixed(2)}</div>
          </div>
          <div className="mt-3 text-sm text-slate-300">{n.whySelected}</div>

          {!!n.alternatives?.length && (
            <div className="mt-4 grid gap-3 md:grid-cols-2">
              {n.alternatives.map((a, i) => (
                <div key={`${n.nodeId}-${i}`} className="rounded-xl border border-white/10 bg-slate-900/70 p-3 text-sm">
                  <div className="font-medium">{a.action}</div>
                  <div className="mt-1 text-slate-400">Score {Number(a.score || 0).toFixed(2)}</div>
                  <div className="mt-2 text-slate-300">{a.reason}</div>
                </div>
              ))}
            </div>
          )}
        </div>
      ))}
    </div>
  );
}

function CompliancePanel({ summary, projectedSummary, nodeIntel = [] }) {
  const maxBlastNode = [...nodeIntel].sort((a, b) => (b.blastRadius || 0) - (a.blastRadius || 0))[0];

  return (
    <div className="grid gap-4 md:grid-cols-[360px_1fr]">
      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="text-lg font-semibold">Policy Posture</div>
        <div className="mt-4 space-y-3">
          <Metric label="Current Compliance" value={summary?.complianceScore} />
          <Metric label="Projected Compliance" value={projectedSummary?.projectedComplianceScore} />
          <Metric label="Current Cost Risk" value={summary?.costRiskScore} />
          <Metric label="Projected Cost Risk" value={projectedSummary?.projectedCostRiskScore} />
          <Metric label="Highest Blast Radius Node" value={maxBlastNode?.nodeId || "—"} />
          <Metric label="Blast Radius" value={maxBlastNode?.blastRadius || "—"} />
        </div>
      </div>

      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="text-lg font-semibold">Blast Radius View</div>
        <div className="mt-4 grid gap-3 md:grid-cols-2">
          {nodeIntel.map((n) => (
            <div key={n.nodeId} className="rounded-xl border border-white/10 bg-slate-900/70 p-3">
              <div className="font-medium">{n.nodeId}</div>
              <div className="mt-1 text-xs text-slate-500">
                Blast Radius {n.blastRadius} · Risk Influence {Number(n.riskInfluence || 0).toFixed(2)}
              </div>
              <div className="mt-2 text-sm text-slate-300">{n.why}</div>
              {!!n.affectedNodes?.length && (
                <div className="mt-3 flex flex-wrap gap-2">
                  {n.affectedNodes.map((v) => (
                    <span
                      key={`${n.nodeId}-${v}`}
                      className="rounded-full border border-white/10 bg-slate-950/60 px-2.5 py-1 text-[11px] text-slate-300"
                    >
                      {v}
                    </span>
                  ))}
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

function Metric({ label, value }) {
  return (
    <div className="rounded-xl border border-white/10 bg-slate-900/70 p-3">
      <div className="text-xs uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 text-lg font-semibold">{value ?? "—"}</div>
    </div>
  );
}

export default function AnalysisWorkspacePage({
  activeTab,
  setActiveTab,
  state,
  session,
  onBack,
  onSelectNode,
  onSelectPath,
  selectedPathIndex,
  selectedNode,
  selectedRecommendation,
  selectedForecast,
  selectedAttackPath,
  onRefreshState,
}) {
  return (
    <div className="grid h-full grid-rows-[auto_1fr] gap-3">
      <div className="flex items-center justify-between gap-3">
        <div className="min-w-0 flex-1 overflow-hidden">
          <WorkspaceTabs activeTab={activeTab} setActiveTab={setActiveTab} />
        </div>
        <button
          onClick={onBack}
          className="shrink-0 rounded-xl border border-white/10 bg-slate-950/60 px-4 py-2 text-sm"
        >
          Back to Control Plane
        </button>
      </div>

      <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        {activeTab === "graph" && (
          <div className="grid h-full grid-rows-[64px_1fr] gap-3">
            <GraphActivityPanel
              metrics={{
                attackPathCount: state.attackMetrics?.pathCount ?? state.attackPaths.length,
                publicExposureCount: state.summary?.publicExposureCount ?? 0,
                maxBlastRadius: Math.max(...(state.nodeIntel || []).map((n) => n.blastRadius || 0), 0),
              }}
              selectedAttackPath={selectedAttackPath}
              onOpenWorkspace={setActiveTab}
            />

            <div className="min-h-0">
              <div className="grid h-full grid-cols-[minmax(0,1fr)_300px] gap-3">
                <div className="min-h-0">
                  <GraphPanel
                    nodes={state.nodes}
                    edges={state.edges}
                    nodeIntel={state.nodeIntel}
                    attackPathCount={state.attackMetrics?.pathCount ?? state.attackPaths.length}
                    selectedAttackPath={selectedAttackPath}
                    onSelectNode={(node) => onSelectNode(node?.id)}
                  />
                </div>

                <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-3">
                  <NodeDrawer
                    node={selectedNode}
                    recommendation={selectedRecommendation}
                    forecast={selectedForecast}
                    nodeIntel={state.nodeIntel}
                  />
                </div>
              </div>
            </div>
          </div>
        )}

        {activeTab === "governance" && (
          <RecommendationList
            recommendations={state.recommendations}
            onSelect={(rec) => onSelectNode(rec.nodeId)}
          />
        )}

        {activeTab === "explainability" && (
          <ExplainabilityList
            explanations={state.explanations}
            onSelect={(nodeId) => onSelectNode(nodeId)}
          />
        )}

        {activeTab === "billshock" && (
          <ForecastPanel forecasts={state.forecasts} onSelect={onSelectNode} />
        )}

                {activeTab === "gitops" && (
          <GitOpsPanel
            gitops={state.gitops}
            session={session}
            scenario={state.scenario}
            seed={state.seed}
            onRefreshState={onRefreshState}
          />
        )}

        {activeTab === "feedback" && <FeedbackPanel feedback={state.feedback} />}

        {activeTab === "carbon" && (
          <CarbonPanel summary={state.summary} projectedSummary={state.projectedSummary} />
        )}

        {activeTab === "negotiation" && (
          <NegotiationPanel negotiations={state.negotiations} />
        )}

        {activeTab === "compliance" && (
          <CompliancePanel
            summary={state.summary}
            projectedSummary={state.projectedSummary}
            nodeIntel={state.nodeIntel}
          />
        )}

        {activeTab === "attackpaths" && (
          <AttackPathsPanel
            attackPaths={state.attackPaths}
            onSelectPath={onSelectPath}
            selectedIndex={selectedPathIndex}
          />
        )}

        {activeTab === "propagation" && (
          <PropagationPanel
            nodeIntel={state.nodeIntel}
            edgeInfluence={state.edgeInfluence || []}
          />
        )}
      </div>
    </div>
  );
}