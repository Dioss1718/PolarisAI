const { HIGH_RISK_EXPOSURE } = require("../config/rules");

function applyRiskMutation(data) {
  data.nodes.forEach(node => {
    node.compliance_flags = node.compliance_flags || [];
    node.activity_logs = node.activity_logs || [];

    if (node.exposure === "PUBLIC") {
      if (!node.compliance_flags.includes(HIGH_RISK_EXPOSURE)) {
        node.compliance_flags.push(HIGH_RISK_EXPOSURE);
      }

      if (!node.activity_logs.includes("Public exposure risk detected")) {
        node.activity_logs.push("Public exposure risk detected");
      }
    }
  });

  return data;
}

module.exports = applyRiskMutation;