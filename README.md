# PolarisAI — Autonomous Cloud Governance Engine

PolarisAI is a graph-driven autonomous governance system that analyzes multi-cloud infrastructure, detects security risks, optimizes cost, enforces policies, and executes safe remediation through GitOps. It transforms cloud governance from reactive monitoring into a fully autonomous, decision-making pipeline powered by real graph intelligence, multi-objective optimization, explainable AI, and a self-improving feedback loop.

---

## Prerequisites

Before running PolarisAI, ensure the following are installed on your machine:

- Go 1.21 or higher
- Node.js 18 or higher
- Python 3.10 or higher
- pip (Python package manager)
- A valid Groq API key (free at https://console.groq.com)

---

## Environment Setup

Create a `.env` file in the root folder and add:

```
GROQ_API_KEY=your_key
GITHUB_TOKEN=your_token
GITHUB_REPO=your_repo
AI_ENGINE_URL=http://127.0.0.1:8000
SIMULATION_API_URL=http://127.0.0.1:7000
```

---

## One Manual Step Required

The AI knowledge base must be ingested once before the AI engine can explain decisions. This does not need to be repeated on subsequent runs.

File location: `cloud/ai-engine/ingest.py`

```bash
cd ai-engine
python ingest.py
```

This embeds the cloud governance knowledge base (SLA documents, compliance benchmarks, security policies, architecture guides) into the local ChromaDB vector store. It takes 1 to 3 minutes on first run.

---

## One-Command Startup

After ingest is complete, run one of the following scripts from the project root. They start all services in the correct order automatically.

**On Windows (PowerShell)** — file: `cloud/scripts/run_demo.ps1`

```powershell
.\scripts\run_demo.ps1
```

If blocked by execution policy, run this first:

```powershell
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
```

**On Linux or macOS (Bash)** — file: `cloud/scripts/run_demo.sh`

```bash
chmod +x scripts/run_demo.sh
./scripts/run_demo.sh
```

Do not close any terminal window that opens. All six services must remain running.

---

## Service Ports

| Service | URL |
|---|---|
| Frontend (React + Vite) | http://127.0.0.1:5173 |
| Go Backend (Governance API) | http://127.0.0.1:8080 |
| Python AI Engine (Explainability + Copilot) | http://localhost:8000 |
| Python Forecast Service (Bill Shock) | http://localhost:5050 |
| Synthetic Engine (Simulation) | http://localhost:7000 |

Open the application at: **http://127.0.0.1:5173**

Synthetic Engine simulation endpoint: **http://localhost:7000/simulate/run**

---

## Login Credentials

| Role | Employee ID | Password | Access |
|---|---|---|---|
| Admin | AT001 | admin123 | Full access including Simulation Studio and GitOps merge |
| DevOps | AT002 | devops123 | Governance, graph, simulation, GitOps view, bill shock |
| Security | AT003 | security123 | SECURE actions only, graph, explainability |

---

## Running the Governance Pipeline

Once logged in, select a scenario from the dropdown (FULL_CHAOS recommended), set seed to 42, and click Run Governance. The pipeline executes all stages and populates the dashboard with real computed data.

If your role is Admin or DevOps, the Simulation Studio panel on the main dashboard lets you inject a custom JSON graph directly into the pipeline, bypassing the synthetic engine.

---

## How It Works

The governance pipeline runs twelve stages in sequence. Each stage feeds into the next with no hardcoded values anywhere in the pipeline.

- **Graph Build** — Constructs a multi-cloud dependency graph. Nodes are AWS, Azure, and GCP resources. Edges are IAM, network, and storage access relationships.
- **Risk Modeling** — Dijkstra-based risk propagation across the graph using exposure, criticality, compliance flags, and connectivity.
- **Security Sentinel** — BFS and DFS traversal from public entry points to sensitive assets, computing real reachable attack paths.
- **Cost Optimizer** — Analyzes CPU utilization and spend. Flags underutilized and overexposed resources.
- **Candidate Generator** — Fuses risk and cost signals into ranked remediation candidates per node.
- **Action Generator** — Produces TERMINATE, DOWNSIZE, or SECURE action variants (TERMINATE_SAFE, DOWNSIZE_MEDIUM, SECURE_PATCH, SECURE_RESTRICT, and others).
- **Pareto Optimizer** — Multi-objective optimization across all actions. Balances risk reduction against cost delta using feedback-learned weights.
- **Policy Validator** — Enforces SLA, compliance, production safety, and blast radius constraints. Outputs APPROVED, MODIFIED, or REJECTED per action.
- **AI Explainability** — RAG pipeline retrieves grounded evidence from ChromaDB (CIS benchmarks, SLA docs, security policies) and generates human-readable reasoning via Llama 3.1 on Groq.
- **Forecast Engine** — Predicts 30-day and 90-day cost trends using pre-trained ML models. Detects bill shock conditions.
- **GitOps Execution** — Converts approved decisions into branch-isolated GitHub pull requests with rollback information.
- **Feedback Learning Loop** — Records outcomes and updates risk weight, cost weight, and penalty coefficient for the Pareto optimizer, making the system self-improving over time.

After the pipeline runs, click Workspace in the header to open the Analysis Workspace. The tabs available are: Graph Workspace (interactive node graph with attack path highlighting), Attack Paths, Governance Actions, Explainability, Bill Shock Watch, GitOps, Carbon Intelligence, Compliance and Blast Radius, Risk Propagation, Negotiation and Tradeoffs, and Adaptive Feedback.

---

## Copilot

Click the Copilot button at the bottom-right to open the AI copilot, grounded on the current pipeline state and powered by Llama 3.1 via Groq. Example questions:

- Which node has the highest risk?
- Are there any bill shock risks?
- Explain the attack paths
- What actions were approved?

---

## Technology Stack

| Layer | Technology |
|---|---|
| Backend Orchestrator | Go 1.21 |
| Frontend | React 18, Vite, Tailwind CSS |
| Synthetic Engine | Node.js, Express |
| AI Engine | Python, FastAPI, Groq API (Llama 3.1 8B) |
| RAG Pipeline | ChromaDB, sentence-transformers (all-MiniLM-L6-v2) |
| Forecast Engine | Python, scikit-learn (.pkl models) |
| Graph Visualization | React Flow |
| Authentication | bcrypt token-based sessions |
| Optimization | Pareto multi-objective optimizer (Go) |

---

## Suggested Demo Flow

1. Start all services using the run script
2. Open http://127.0.0.1:5173 and log in as AT001 / admin123
3. Select FULL_CHAOS, seed 42, click Run Governance
4. Observe KPI cards, Before/After panels, and the pipeline stage ticker
5. Open Graph Workspace, click a node to inspect its governance decision
6. Switch to Attack Paths and select a path to highlight it on the graph
7. Open Governance Actions to review APPROVED and MODIFIED recommendations
8. Open Explainability to read AI-generated reasoning for each decision
9. Open the Copilot and ask: which node has the highest risk?
10. Open Bill Shock Watch for 30-day and 90-day cost forecasts
11. Open Carbon Intelligence to see emission reduction after remediation
12. Switch to Negotiation and Tradeoffs to inspect Pareto decision traces