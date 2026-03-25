function simulateTrafficSpike(data) {
  data.nodes.forEach(node => {
    node.metrics = node.metrics || {};
    node.activity_logs = node.activity_logs || [];

    const spike = Math.random() > 0.7;

    if (spike) {
      node.metrics.network_traffic = "SPIKE";

      if (!node.activity_logs.includes("Traffic spike detected")) {
        node.activity_logs.push("Traffic spike detected");
      }
    }
  });

  return data;
}

module.exports = simulateTrafficSpike;