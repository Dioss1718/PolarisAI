const express = require("express");
const router = express.Router();
const { runSimulation } = require("../services/simulationService");

router.get("/", (req, res) => {
  res.json({
    message: "Synthetic Engine Running",
    endpoint: "/simulate/run",
    supported_scenarios: ["FULL_CHAOS", "SECURITY_BREACH", "COST_SPIKE", "POLICY_DRIFT"],
    usage: "/simulate/run?scenario=FULL_CHAOS&seed=42"
  });
});

router.get("/run", (req, res) => {
  const scenario = req.query.scenario || "FULL_CHAOS";
  const seed = Number(req.query.seed || 42);

  const result = runSimulation({ scenario, seed });
  res.json(result);
});

module.exports = router;