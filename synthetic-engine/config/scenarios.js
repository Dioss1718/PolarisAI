module.exports = {
  FULL_CHAOS: {
    applyCost: true,
    applyRisk: true,
    applyTraffic: true,
    applyTopology: true,
    applyExposure: true,
    applyIAM: true,
    applyAttack: true
  },
  SECURITY_BREACH: {
    applyCost: false,
    applyRisk: true,
    applyTraffic: true,
    applyTopology: true,
    applyExposure: true,
    applyIAM: true,
    applyAttack: true
  },
  COST_SPIKE: {
    applyCost: true,
    applyRisk: false,
    applyTraffic: false,
    applyTopology: false,
    applyExposure: false,
    applyIAM: false,
    applyAttack: false
  },
  POLICY_DRIFT: {
    applyCost: false,
    applyRisk: true,
    applyTraffic: false,
    applyTopology: false,
    applyExposure: true,
    applyIAM: true,
    applyAttack: false
  }
};