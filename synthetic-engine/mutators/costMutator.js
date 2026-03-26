const { COST_SPIKE_FACTOR, LOW_UTIL_THRESHOLD } = require("../config/rules");
const { logEvent } = require("../utils/eventLogger");
const { computeSeverity } = require("../utils/severity");

function applyCostMutation(data, random) {
  data.nodes.forEach(node => {
    node.activity_logs = node.activity_logs || [];

    if (node.cost > 0 && node.utilization < LOW_UTIL_THRESHOLD) {
      const oldCost = node.cost;
      node.cost = Math.round(node.cost * COST_SPIKE_FACTOR);

      const message = "Cost anomaly: idle resource detected";
      if (!node.activity_logs.includes(message)) {
        node.activity_logs.push(message);
      }

      logEvent(data, {
        type: "COST_ANOMALY",
        node_id: node.id,
        old_cost: oldCost,
        new_cost: node.cost,
        reason: "Low utilization resource cost spike",
        severity: computeSeverity(node)
      });
    }
  });

  return data;
}

module.exports = applyCostMutation;