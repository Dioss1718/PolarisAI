function mutateIAM(data) {
  data.nodes.forEach(node => {
    if (node.type === "IAM_ROLE") {
      const rand = Math.random();

      // Privilege escalation
      if (rand > 0.7 && !node.compliance_flags.includes("ADMIN_ACCESS")) {
        node.compliance_flags.push("ADMIN_ACCESS");
        node.activity_logs.push("Privilege escalation detected");
      }

      // Privilege reduction
      if (rand < 0.2 && node.compliance_flags.includes("ADMIN_ACCESS")) {
        node.compliance_flags = node.compliance_flags.filter(
          f => f !== "ADMIN_ACCESS"
        );
        node.activity_logs.push("Privileges reduced");
      }
    }
  });

  return data;
}

module.exports = mutateIAM;