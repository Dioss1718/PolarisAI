const { logEvent } = require("../utils/eventLogger");

function mutateTopology(data, random) {
  if (!data.nodes || data.nodes.length < 2) return data;

  data.edges = data.edges || [];
  data.logs = data.logs || {};
  data.logs.api_logs = data.logs.api_logs || [];

  const roll = random();

  if (roll > 0.6) {
    const from = data.nodes[Math.floor(random() * data.nodes.length)];
    const to = data.nodes[Math.floor(random() * data.nodes.length)];

    if (from.id !== to.id) {
      const exists = data.edges.some(e => e.from === from.id && e.to === to.id);

      if (!exists) {
        const edge = {
          from: from.id,
          to: to.id,
          type: "DYNAMIC_LINK",
          weight: Math.floor(random() * 5) + 1
        };

        data.edges.push(edge);
        data.logs.api_logs.push(`Dynamic link created: ${from.id} -> ${to.id}`);

        logEvent(data, {
          type: "TOPOLOGY_DRIFT",
          edge,
          reason: "Unexpected dependency introduced",
          severity: "MEDIUM"
        });
      }
    }
  }

  if (roll < 0.3 && data.edges.length > 0) {
    const index = Math.floor(random() * data.edges.length);
    const removed = data.edges.splice(index, 1)[0];
    data.logs.api_logs.push(`Edge removed: ${removed.from} -> ${removed.to}`);

    logEvent(data, {
      type: "TOPOLOGY_REMOVAL",
      edge: removed,
      reason: "Dependency removed",
      severity: "LOW"
    });
  }

  return data;
}

module.exports = mutateTopology;