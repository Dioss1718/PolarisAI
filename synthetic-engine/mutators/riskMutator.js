const { HIGH_RISK_EXPOSURE } = require("../config/rules");
const { logEvent } = require("../utils/eventLogger");
const { computeSeverity } = require("../utils/severity");

function applyRiskMutation(data, random) {
  data.nodes.forEach(node => {
    node.compliance_flags = node.compliance_flags || [];
    node.activity_logs = node.activity_logs || [];

    if (node.exposure === "PUBLIC") {
      if (!node.compliance_flags.includes(HIGH_RISK_EXPOSURE)) {
        node.compliance_flags.push(HIGH_RISK_EXPOSURE);
      }

      const message = "Public exposure risk detected";
      if (!node.activity_logs.includes(message)) {
        node.activity_logs.push(message);
      }

      logEvent(data, {
        type: "SECURITY_EXPOSURE",
        node_id: node.id,
        reason: "Public exposure increases attack surface",
        severity: computeSeverity(node)
      });
    }
  });

  return data;
}

module.exports = applyRiskMutation;