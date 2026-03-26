const { logEvent } = require("../utils/eventLogger");
const { computeSeverity } = require("../utils/severity");

function simulateTrafficSpike(data, random) {
  data.nodes.forEach(node => {
    node.metrics = node.metrics || {};
    node.activity_logs = node.activity_logs || [];

    const spike = random() > 0.7;

    if (spike) {
      node.metrics.network_traffic = "SPIKE";

      const message = "Traffic spike detected";
      if (!node.activity_logs.includes(message)) {
        node.activity_logs.push(message);
      }

      logEvent(data, {
        type: "TRAFFIC_SPIKE",
        node_id: node.id,
        reason: "Sudden network traffic increase",
        severity: computeSeverity(node)
      });
    }
  });

  return data;
}

module.exports = simulateTrafficSpike;