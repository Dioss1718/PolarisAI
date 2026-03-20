const { COST_SPIKE_FACTOR, LOW_UTIL_THRESHOLD } = require("../config/rules");

function applyCostMutation(data) {
  data.nodes.forEach(node => {
    node.activity_logs = node.activity_logs || [];

    if (node.utilization < LOW_UTIL_THRESHOLD) {
      node.cost = Math.round(node.cost * COST_SPIKE_FACTOR);

      if (!node.activity_logs.includes("Cost anomaly: idle resource detected")) {
        node.activity_logs.push("Cost anomaly: idle resource detected");
      }
    }
  });

  return data;
}

module.exports = applyCostMutation;