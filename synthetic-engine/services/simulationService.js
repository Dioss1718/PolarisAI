const fs = require("fs");
const path = require("path");

const deepClone = require("../utils/deepClone");

const applyCostMutation = require("../mutators/costMutator");
const applyRiskMutation = require("../mutators/riskMutator");
const injectAttackPath = require("../mutators/attackMutator");
const simulateTrafficSpike = require("../mutators/trafficMutator");

// NEW IMPORTS
const mutateTopology = require("../mutators/topologyMutator");
const mutateExposure = require("../mutators/exposureMutator");
const mutateIAM = require("../mutators/iamMutator");

const datasetPath = path.join(__dirname, "../../Data/cloud_env.json");

function runSimulation() {
  try {
    const rawData = JSON.parse(fs.readFileSync(datasetPath, "utf-8"));

    let data = deepClone(rawData);

    data.timestamp = new Date().toISOString();

    data.logs = data.logs || {};
    data.logs.api_logs = data.logs.api_logs || [];

    // ------------------------
    // CORE MUTATIONS
    // ------------------------
    data = applyCostMutation(data);
    data = applyRiskMutation(data);
    data = simulateTrafficSpike(data);

    // ------------------------
    // NEW ADVANCED MUTATIONS
    // ------------------------
    data = mutateTopology(data);
    data = mutateExposure(data);
    data = mutateIAM(data);

    // ------------------------
    // ATTACK INJECTION
    // ------------------------
    data = injectAttackPath(data);

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