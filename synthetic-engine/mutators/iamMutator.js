const { logEvent } = require("../utils/eventLogger");
const { computeSeverity } = require("../utils/severity");

function mutateIAM(data, random) {
  data.nodes.forEach(node => {
    node.compliance_flags = node.compliance_flags || [];
    node.activity_logs = node.activity_logs || [];

    if (node.type === "IAM_ROLE") {
      const rand = random();

      if (rand > 0.7 && !node.compliance_flags.includes("ADMIN_ACCESS")) {
        node.compliance_flags.push("ADMIN_ACCESS");
        node.activity_logs.push("Privilege escalation detected");

        logEvent(data, {
          type: "IAM_ESCALATION",
          node_id: node.id,
          reason: "Role gained admin privileges",
          severity: computeSeverity(node)
        });
      }

      if (rand < 0.2 && node.compliance_flags.includes("ADMIN_ACCESS")) {
        node.compliance_flags = node.compliance_flags.filter(f => f !== "ADMIN_ACCESS");
        node.activity_logs.push("Privileges reduced");

        logEvent(data, {
          type: "IAM_REDUCTION",
          node_id: node.id,
          reason: "Role privileges reduced",
          severity: "LOW"
        });
      }
    }
  });

  return data;
}

module.exports = mutateIAM;