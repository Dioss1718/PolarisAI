function mutateTopology(data) {
  if (!data.nodes || data.nodes.length < 2) return data;

  data.edges = data.edges || [];

  const random = Math.random();

  // Add new unexpected connection
  if (random > 0.6) {
    const from = data.nodes[Math.floor(Math.random() * data.nodes.length)];
    const to = data.nodes[Math.floor(Math.random() * data.nodes.length)];

    if (from.id !== to.id) {
      const exists = data.edges.some(
        e => e.from === from.id && e.to === to.id
      );

      if (!exists) {
        data.edges.push({
          from: from.id,
          to: to.id,
          type: "DYNAMIC_LINK",
          weight: Math.floor(Math.random() * 5) + 1
        });

        data.logs.api_logs.push(`Dynamic link created: ${from.id} -> ${to.id}`);
      }
    }
  }

  // Randomly remove weak edges
  if (random < 0.3 && data.edges.length > 0) {
    const index = Math.floor(Math.random() * data.edges.length);
    const removed = data.edges.splice(index, 1);

    data.logs.api_logs.push(`Edge removed: ${removed[0].from} -> ${removed[0].to}`);
  }

  return data;
}

module.exports = mutateTopology;