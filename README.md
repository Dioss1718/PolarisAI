# PolarisAI - Autonomous Cloud Governance Engine

PolarisAI is a graph-driven system that analyzes multi-cloud infrastructure, detects risks, optimizes cost, enforces policies, and executes safe remediation through GitOps. It transforms cloud governance from reactive monitoring into an autonomous, decision-making pipeline.

---

## Architecture Overview

PolarisAI operates as a structured governance pipeline:

Simulation Engine  
→ Graph Engine (multi-cloud topology)  
→ Security and Cost Analysis  
→ Candidate and Action Generation  
→ Pareto Optimization (risk vs cost)  
→ Policy Validation (SLA, compliance)  
→ AI Explainability  
→ Forecast Engine (cost prediction)  
→ GitOps Execution (pull request creation)  
→ Feedback Learning Loop  

---

## How to Run Locally

Backend (Go)

cd backend  
go run main.go  

Frontend (React + Vite)

cd frontend  
npm install  
npm run dev  

Open in browser:

http://127.0.0.1:5173  

---

## Environment Variables

Create a .env file:

GITHUB_TOKEN=your_token_here  
GITHUB_OWNER=your_username  
GITHUB_REPO=your_repo  

---

## Key Features

Unified Cloud Graph  
- Represents AWS, Azure, and GCP resources in a single topology  
- Captures dependencies and relationships between resources  
- Uses risk-based coloring for visibility  

Security and Risk Engine  
- Identifies exposed resources and attack paths  
- Computes node-level risk scores  

Cost Optimization  
- Detects underutilized resources  
- Recommends cost-efficient actions  

Autonomous Decision Engine  
- Generates multiple remediation strategies  
- Evaluates trade-offs between cost and risk  
- Selects optimal actions using Pareto optimization  

Policy Validation  
- Enforces SLA, compliance, and security constraints  
- Outputs approved, modified, or rejected actions  

AI Explainability  
- Generates human-readable reasoning for decisions  
- Improves transparency and auditability  

Forecast Engine  
- Predicts future cost trends (30-day and 90-day)  
- Detects potential bill shock scenarios  

GitOps Execution  
- Converts decisions into infrastructure changes  
- Automatically creates GitHub pull requests  
- Enables safe and reviewable execution  

Feedback Learning Loop  
- Learns from past decisions and outcomes  
- Adjusts weighting between risk, cost, and penalty  

Simulation Engine  
- Supports predefined scenarios such as chaos and cost spikes  
- Allows manual JSON-based infrastructure simulation  

Operator Command Space  
- Enables query-driven interaction with system state  
- Provides insights into risks, costs, and decisions  

---

## Summary

PolarisAI combines graph intelligence, multi-objective optimization, explainable AI, and GitOps execution to create a complete autonomous cloud governance system. It enables organizations to move from reactive monitoring to proactive, policy-aware, and self-improving infrastructure management.