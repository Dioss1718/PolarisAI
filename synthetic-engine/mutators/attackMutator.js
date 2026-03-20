function injectAttackPath(data) {
  if (!data.nodes || data.nodes.length < 2) return data;

  const attacker = data.nodes[0];
  const target = data.nodes[data.nodes.length - 1];

  data.edges = data.edges || [];
  data.logs = data.logs || {};
  data.logs.api_logs = data.logs.api_logs || [];

  const exists = data.edges.some(
    e => e.type === "SIMULATED_ATTACK" && e.from === attacker.id && e.to === target.id
  );

  if (!exists) {
    data.edges.push({
      from: attacker.id,
      to: target.id,
      type: "SIMULATED_ATTACK",
      weight: 5
    });

    data.logs.api_logs.push("Simulated attack path injected");
  }

  return data;
}

module.exports = injectAttackPath;