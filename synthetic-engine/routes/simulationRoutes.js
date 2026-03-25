const express = require("express");
const router = express.Router();
const { runSimulation } = require("../services/simulationService");

// Health route (for demo clarity)
router.get("/", (req, res) => {
  res.json({
    message: "Synthetic Engine Running",
    endpoint: "/simulate/run"
  });
});

// Simulation route
router.get("/run", (req, res) => {
  const result = runSimulation();
  res.json(result);
});

module.exports = router;