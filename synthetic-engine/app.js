const express = require("express");
const cors = require("cors");

const app = express();
app.use(cors());
app.use(express.json());

const simulationRoutes = require("./routes/simulationRoutes");

app.use("/simulate", simulationRoutes);

const PORT = 7000;

app.listen(PORT, () => {
  console.log(`Synthetic Mutation Engine running on port ${PORT}`);
});