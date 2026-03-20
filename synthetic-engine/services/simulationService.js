const fs = require("fs");
const path = require("path");

const deepClone = require("../utils/deepClone");

const applyCostMutation = require("../mutators/costMutator");
const applyRiskMutation = require("../mutators/riskMutator");
const injectAttackPath = require("../mutators/attackMutator");
const simulateTrafficSpike = require("../mutators/trafficMutator");

const datasetPath = path.join(__dirname, "../../Data/cloud_env.json");

function runSimulation() {
  try {
    const rawData = JSON.parse(fs.readFileSync(datasetPath, "utf-8"));

    let data = deepClone(rawData);

    // Add timestamp 
    data.timestamp = new Date().toISOString();

    // Apply mutations
    data = applyCostMutation(data);
    data = applyRiskMutation(data);
    data = simulateTrafficSpike(data);
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