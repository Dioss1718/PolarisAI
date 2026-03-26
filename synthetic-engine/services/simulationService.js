const fs = require("fs");
const path = require("path");

const scenarios = require("../config/scenarios");
const deepClone = require("../utils/deepClone");
const { createSeededRandom } = require("../utils/seededRandom");

const applyCostMutation = require("../mutators/costMutator");
const applyRiskMutation = require("../mutators/riskMutator");
const injectAttackPath = require("../mutators/attackMutator");
const simulateTrafficSpike = require("../mutators/trafficMutator");
const mutateTopology = require("../mutators/topologyMutator");
const mutateExposure = require("../mutators/exposureMutator");
const mutateIAM = require("../mutators/iamMutator");

const datasetPath = path.join(__dirname, "../../Data/cloud_env.json");

function initializeExpectedIssues(data) {
  data.expected_issues = [];
  data.simulation_metadata = data.simulation_metadata || {};
  return data;
}

function deriveExpectedIssues(data) {
  data.nodes.forEach(node => {
    if (node.exposure === "PUBLIC" && node.criticality >= 8) {
      data.expected_issues.push({
        node_id: node.id,
        issue_type: "HIGH_RISK_EXPOSURE",
        severity: "HIGH"
      });
    }

    if (node.cost > 100 && node.utilization <= 20) {
      data.expected_issues.push({
        node_id: node.id,
        issue_type: "COST_ANOMALY",
        severity: "HIGH"
      });
    }

    if (node.type === "IAM_ROLE" && (node.compliance_flags || []).includes("ADMIN_ACCESS")) {
      data.expected_issues.push({
        node_id: node.id,
        issue_type: "IAM_ESCALATION",
        severity: "HIGH"
      });
    }
  });

  return data;
}

function runSimulation(options = {}) {
  try {
    const scenario = options.scenario || "FULL_CHAOS";
    const seed = Number.isFinite(options.seed) ? options.seed : 42;
    const scenarioConfig = scenarios[scenario] || scenarios.FULL_CHAOS;
    const random = createSeededRandom(seed);

    const rawData = JSON.parse(fs.readFileSync(datasetPath, "utf-8"));
    let data = deepClone(rawData);

    data.timestamp = new Date().toISOString();
    data.logs = data.logs || {};
    data.logs.api_logs = data.logs.api_logs || [];
    data = initializeExpectedIssues(data);

    if (scenarioConfig.applyCost) {
      data = applyCostMutation(data, random);
    }
    if (scenarioConfig.applyRisk) {
      data = applyRiskMutation(data, random);
    }
    if (scenarioConfig.applyTraffic) {
      data = simulateTrafficSpike(data, random);
    }
    if (scenarioConfig.applyTopology) {
      data = mutateTopology(data, random);
    }
    if (scenarioConfig.applyExposure) {
      data = mutateExposure(data, random);
    }
    if (scenarioConfig.applyIAM) {
      data = mutateIAM(data, random);
    }
    if (scenarioConfig.applyAttack) {
      data = injectAttackPath(data, random);
    }

    data = deriveExpectedIssues(data);

    data.simulation_metadata = {
      scenario,
      seed,
      mutations_applied: Object.entries(scenarioConfig)
        .filter(([, enabled]) => enabled)
        .map(([k]) => k)
    };

    return data;
  } catch (error) {
    console.error("Simulation Error:", error.message);
    return {
      error: "Failed to run simulation",
      details: error.message
    };
  }
}

module.exports = { runSimulation };