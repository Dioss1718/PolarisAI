function num(v) {
  return Number.isFinite(Number(v)) ? Number(v) : 0;
}

export function deriveMetrics(state) {
  const summary = state?.summary || {};
  const projected = state?.projectedSummary || {};
  const nodeIntel = Array.isArray(state?.nodeIntel) ? state.nodeIntel : [];

  return {
    totalNodes: num(summary.totalNodes),
    totalEdges: num(summary.totalEdges),
    attackPathCount: num(summary.attackPathCount),
    avgAttackPathLength: num(summary.avgAttackPathLength),
    reachableNodes: num(summary.reachableNodes),
    highRiskCount: num(summary.highRiskCount),
    publicExposureCount: num(summary.publicExposureCount),
    approvedCount: num(summary.approvedCount),
    modifiedCount: num(summary.modifiedCount),
    rejectedCount: num(summary.rejectedCount),
    urgentCount: num(summary.urgentCount),
    billShockCount: num(summary.billShockCount),
    currentTotalCost: num(summary.currentTotalCost),
    forecast30Total: num(summary.forecast30Total),
    forecast90Total: num(summary.forecast90Total),
    averageRisk: num(summary.averageRisk),
    totalRisk: num(summary.totalRisk),
    currentCarbonTotal: num(summary.currentCarbonTotal),
    complianceScore: num(summary.complianceScore),
    costRiskScore: num(summary.costRiskScore),

    projectedTotalCost: num(projected.projectedTotalCost),
    projectedAttackPathCount: num(projected.projectedAttackPathCount),
    projectedPublicExposureCount: num(projected.projectedPublicExposureCount),
    projectedAverageRisk: num(projected.projectedAverageRisk),
    projectedCarbonTotal: num(projected.projectedCarbonTotal),
    carbonReductionPct: num(projected.carbonReductionPct),
    greenScore: num(projected.greenScore),
    projectedComplianceScore: num(projected.projectedComplianceScore),
    projectedCostRiskScore: num(projected.projectedCostRiskScore),
    projectedRiskReductionPct: num(projected.projectedRiskReductionPct),

    maxBlastRadius: nodeIntel.reduce((m, n) => Math.max(m, Number(n?.blastRadius || 0)), 0),
  };
}

export function deriveAlerts(state) {
  if (Array.isArray(state?.alerts) && state.alerts.length) return state.alerts;

  const m = deriveMetrics(state);
  const alerts = [];

  if (m.urgentCount > 0) {
    alerts.push({
      severity: "critical",
      title: "Urgent governance actions pending",
      metric: `${m.urgentCount} actions`,
      reason: "Approved or modified actions with elevated residual risk are present.",
      workspaceTab: "governance",
    });
  }

  if (m.billShockCount > 0) {
    alerts.push({
      severity: "high",
      title: "Bill shock watch active",
      metric: `${m.billShockCount} nodes`,
      reason: "Forecast engine flagged rising spend pressure.",
      workspaceTab: "billshock",
    });
  }

  if (m.attackPathCount > 0) {
    alerts.push({
      severity: "medium",
      title: "Attack paths discovered",
      metric: `${m.attackPathCount} paths`,
      reason: "Reachable paths to sensitive assets exist in the graph.",
      workspaceTab: "attackpaths",
    });
  }

  return alerts;
}