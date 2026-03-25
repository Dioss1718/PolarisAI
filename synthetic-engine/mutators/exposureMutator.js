function mutateExposure(data) {
  data.nodes.forEach(node => {
    const rand = Math.random();

    // Simulate accidental exposure
    if (rand > 0.75 && node.exposure !== "PUBLIC") {
      node.exposure = "PUBLIC";
      node.activity_logs.push("Misconfiguration: Resource exposed publicly");
    }

    // Simulate fix
    if (rand < 0.1 && node.exposure === "PUBLIC") {
      node.exposure = "PRIVATE";
      node.activity_logs.push("Auto-remediation: Exposure restricted");
    }
  });

  return data;
}

module.exports = mutateExposure;