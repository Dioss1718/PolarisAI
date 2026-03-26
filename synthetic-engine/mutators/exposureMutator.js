const { logEvent } = require("../utils/eventLogger");
const { computeSeverity } = require("../utils/severity");

function mutateExposure(data, random) {
  data.nodes.forEach(node => {
    node.activity_logs = node.activity_logs || [];
    const rand = random();

    if (rand > 0.75 && node.exposure !== "PUBLIC") {
      const oldExposure = node.exposure;
      node.exposure = "PUBLIC";
      node.activity_logs.push("Misconfiguration: Resource exposed publicly");

      logEvent(data, {
        type: "EXPOSURE_DRIFT",
        node_id: node.id,
        old_exposure: oldExposure,
        new_exposure: node.exposure,
        reason: "Accidental public exposure",
        severity: computeSeverity(node)
      });
    }

    if (rand < 0.1 && node.exposure === "PUBLIC") {
      const oldExposure = node.exposure;
      node.exposure = "PRIVATE";
      node.activity_logs.push("Auto-remediation: Exposure restricted");

      logEvent(data, {
        type: "EXPOSURE_REMEDIATION",
        node_id: node.id,
        old_exposure: oldExposure,
        new_exposure: node.exposure,
        reason: "Exposure restriction applied",
        severity: "LOW"
      });
    }
  });

  return data;
}

module.exports = mutateExposure;