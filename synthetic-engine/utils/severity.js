function computeSeverity({ exposure, criticality = 0, cost = 0, utilization = 100 }) {
  if (exposure === "PUBLIC" && criticality >= 8) return "HIGH";
  if (cost >= 150 && utilization <= 20) return "HIGH";
  if (criticality >= 6 || exposure === "PUBLIC") return "MEDIUM";
  return "LOW";
}

module.exports = { computeSeverity };