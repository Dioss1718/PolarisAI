const { logEvent } = require("../utils/eventLogger");

function injectAttackPath(data, random) {
  if (!data.nodes || data.nodes.length < 2) return data;

  data.edges = data.edges || [];
  data.logs = data.logs || {};
  data.logs.api_logs = data.logs.api_logs || [];

  const publicNodes = data.nodes.filter(n => n.exposure === "PUBLIC");
  const sensitiveNodes = data.nodes.filter(
    n =>
      n.type === "DATABASE" ||
      (n.type === "IAM_ROLE" && (n.compliance_flags || []).includes("ADMIN_ACCESS"))
  );

  if (publicNodes.length === 0 || sensitiveNodes.length === 0) return data;

  const attacker = publicNodes[Math.floor(random() * publicNodes.length)];
  const target = sensitiveNodes[Math.floor(random() * sensitiveNodes.length)];

  const exists = data.edges.some(
    e => e.type === "SIMULATED_ATTACK" && e.from === attacker.id && e.to === target.id
  );

  if (!exists) {
    const edge = {
      from: attacker.id,
      to: target.id,
      type: "SIMULATED_ATTACK",
      weight: 5
    };

    data.edges.push(edge);
    data.logs.api_logs.push("Simulated attack path injected");

    logEvent(data, {
      type: "ATTACK_PATH_INJECTED",
      edge,
      reason: "Synthetic attack chain created",
      severity: "HIGH"
    });
  }

  return data;
}

module.exports = injectAttackPath;