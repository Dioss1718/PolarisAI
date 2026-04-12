from typing import Any, Dict, List, Optional

from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

from retriever import retrieve
from prompt import build_prompt
from gitops_prompt import build_gitops_prompt
from llm import call_llm

load_dotenv()

app = FastAPI(title="Polaris AI Engine", version="1.0.0")


class ExplainRequest(BaseModel):
    node_id: str
    action: str
    env: str
    node_type: str
    cost: float
    risk_reduction: float
    sla: float
    security: float
    compliance: float
    blast: float


class ExplainResponse(BaseModel):
    explanation: str
    grounded: bool
    sources: List[str]


class RetrieveRequest(BaseModel):
    query: str
    node_type: str
    action: str


class RetrieveResponse(BaseModel):
    documents: List[str]
    sources: List[str]


class InfraRequest(BaseModel):
    node_id: str
    action: str
    reason: str
    changes: List[str]
    format: str = "terraform"


class InfraResponse(BaseModel):
    code: str
    format: str
    title: Optional[str] = None
    summary: Optional[str] = None
    grounded: bool = True


class CopilotRequest(BaseModel):
    query: str
    state: Dict[str, Any]


class CopilotResponse(BaseModel):
    answer: str
    grounded: bool
    sources: List[str]


@app.get("/health")
def health():
    return {"status": "ok"}


def _safe_list(value):
    return value if isinstance(value, list) else []


def _safe_dict(value):
    return value if isinstance(value, dict) else {}


def _match_node(query: str, nodes: List[Dict[str, Any]]) -> Optional[Dict[str, Any]]:
    q = (query or "").lower().strip()
    if not q:
        return None

    for node in nodes:
        node_id = str(node.get("id", "")).lower()
        label = str(node.get("label", "")).lower()

        if node_id and node_id in q:
            return node
        if label and label in q:
            return node

    return None


def _build_copilot_context(query: str, state: Dict[str, Any]):
    nodes = _safe_list(state.get("nodes"))
    recommendations = _safe_list(state.get("recommendations"))
    forecasts = _safe_list(state.get("forecasts"))
    attack_paths = _safe_list(state.get("attackPaths"))
    summary = _safe_dict(state.get("summary"))
    projected = _safe_dict(state.get("projectedSummary"))

    matched_node = _match_node(query, nodes)

    top_risks = sorted(
        nodes,
        key=lambda n: float(n.get("risk", 0) or 0),
        reverse=True,
    )[:5]

    approved_actions = [
        r for r in recommendations
        if str(r.get("status", "")).upper() in {"APPROVED", "MODIFIED"}
    ][:5]

    bill_shocks = [f for f in forecasts if bool(f.get("billShock"))][:5]

    matched_rec = None
    if matched_node:
        for rec in recommendations:
            if rec.get("nodeId") == matched_node.get("id"):
                matched_rec = rec
                break

    retrieval_query_parts = [
        query,
        f"scenario {state.get('scenario', '')}",
        "cloud governance",
        "security risk",
        "policy validation",
        "cost optimization",
        "remediation",
    ]

    node_type = ""
    action = ""

    if matched_node:
        node_type = str(matched_node.get("type", "") or "")
        retrieval_query_parts.append(f"node {matched_node.get('id', '')}")
        retrieval_query_parts.append(f"node type {node_type}")
        retrieval_query_parts.append(f"cloud {matched_node.get('cloud', '')}")

    if matched_rec:
        action = str(matched_rec.get("finalAction") or matched_rec.get("action") or "")
        retrieval_query_parts.append(f"action {action}")

    retrieval_query = " ".join([part for part in retrieval_query_parts if part])

    top_risk_lines = []
    for n in top_risks:
        top_risk_lines.append(
            f"- {n.get('id')} | type={n.get('type')} | cloud={n.get('cloud')} | env={n.get('environment')} | risk={n.get('risk')}"
        )

    approved_lines = []
    for r in approved_actions:
        approved_lines.append(
            f"- {r.get('nodeId')} | action={r.get('action')} | final={r.get('finalAction')} | status={r.get('status')} | score={r.get('score')}"
        )

    shock_lines = []
    for f in bill_shocks:
        shock_lines.append(
            f"- {f.get('nodeId')} | current={f.get('currentCost')} | f30={f.get('forecast30')} | f90={f.get('forecast90')} | reason={f.get('shockReason')}"
        )

    matched_node_text = "No directly matched node from the question."
    if matched_node:
        matched_node_text = (
            f"Matched node:\n"
            f"- id={matched_node.get('id')}\n"
            f"- label={matched_node.get('label')}\n"
            f"- type={matched_node.get('type')}\n"
            f"- cloud={matched_node.get('cloud')}\n"
            f"- environment={matched_node.get('environment')}\n"
            f"- exposure={matched_node.get('exposure')}\n"
            f"- risk={matched_node.get('risk')}\n"
            f"- finalAction={matched_node.get('finalAction')}\n"
            f"- status={matched_node.get('status')}\n"
        )

        if matched_rec:
            matched_node_text += (
                f"Recommendation:\n"
                f"- action={matched_rec.get('action')}\n"
                f"- finalAction={matched_rec.get('finalAction')}\n"
                f"- reason={matched_rec.get('reason')}\n"
                f"- riskReduction={matched_rec.get('riskReduction')}\n"
                f"- costDelta={matched_rec.get('costDelta')}\n"
            )

    state_context = f"""
CURRENT PIPELINE STATE
Scenario: {state.get('scenario')}
Seed: {state.get('seed')}
GeneratedAt: {state.get('generatedAt')}

SUMMARY
- totalNodes: {summary.get('totalNodes')}
- totalEdges: {summary.get('totalEdges')}
- attackPathCount: {summary.get('attackPathCount')}
- highRiskCount: {summary.get('highRiskCount')}
- publicExposureCount: {summary.get('publicExposureCount')}
- approvedCount: {summary.get('approvedCount')}
- modifiedCount: {summary.get('modifiedCount')}
- rejectedCount: {summary.get('rejectedCount')}
- billShockCount: {summary.get('billShockCount')}
- currentTotalCost: {summary.get('currentTotalCost')}
- forecast30Total: {summary.get('forecast30Total')}
- forecast90Total: {summary.get('forecast90Total')}
- averageRisk: {summary.get('averageRisk')}
- complianceScore: {summary.get('complianceScore')}
- costRiskScore: {summary.get('costRiskScore')}

PROJECTED SUMMARY
- projectedTotalCost: {projected.get('projectedTotalCost')}
- projectedAttackPathCount: {projected.get('projectedAttackPathCount')}
- projectedPublicExposureCount: {projected.get('projectedPublicExposureCount')}
- projectedAverageRisk: {projected.get('projectedAverageRisk')}
- projectedComplianceScore: {projected.get('projectedComplianceScore')}
- projectedCostRiskScore: {projected.get('projectedCostRiskScore')}
- projectedRiskReductionPct: {projected.get('projectedRiskReductionPct')}

TOP RISKY NODES
{chr(10).join(top_risk_lines) if top_risk_lines else "- none"}

APPROVED OR MODIFIED ACTIONS
{chr(10).join(approved_lines) if approved_lines else "- none"}

BILL SHOCK NODES
{chr(10).join(shock_lines) if shock_lines else "- none"}

ATTACK PATH COUNT
- {len(attack_paths)}

{matched_node_text}
""".strip()

    return retrieval_query, node_type, action, state_context


@app.post("/retrieve", response_model=RetrieveResponse)
def retrieve_docs(data: RetrieveRequest):
    try:
        docs, metas = retrieve(data.query, data.node_type, data.action)

        sources = []
        for meta in metas:
            source = meta.get("source")
            category = meta.get("category")
            if source and category:
                sources.append(f"{category}/{source}")

        return RetrieveResponse(
            documents=docs,
            sources=sources
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@app.post("/explain", response_model=ExplainResponse)
def explain(data: ExplainRequest):
    try:
        query = (
            f"node {data.node_id} "
            f"action {data.action} "
            f"resource type {data.node_type} "
            f"environment {data.env} "
            f"sla impact security implications compliance constraints "
            f"cost tradeoff allowed actions terminate downsize secure"
        )

        docs, metas = retrieve(query, data.node_type, data.action)
        prompt = build_prompt(data.model_dump(), docs)
        output = call_llm(prompt)

        if not output or len(output.strip()) < 20:
            raise ValueError("LLM returned weak output")

        forbidden = ["SCALE", "MIGRATE", "REDEPLOY", "RESTART"]
        fixed_output = output
        for word in forbidden:
            fixed_output = fixed_output.replace(word, "SECURE")

        sources = []
        for meta in metas:
            source = meta.get("source")
            category = meta.get("category")
            if source and category:
                sources.append(f"{category}/{source}")

        return ExplainResponse(
            explanation=fixed_output,
            grounded=len(docs) > 0,
            sources=sources
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@app.post("/copilot", response_model=CopilotResponse)
def copilot(data: CopilotRequest):
    try:
        query = (data.query or "").strip()
        if not query:
            raise ValueError("Query is required")

        retrieval_query, node_type, action, state_context = _build_copilot_context(
            query,
            data.state or {}
        )

        docs, metas = retrieve(retrieval_query, node_type, action)
        docs_text = "\n\n".join(docs) if docs else "NO DOCUMENT EVIDENCE RETRIEVED."

        prompt = f"""
You are the PolarisAI cloud governance copilot.

You must answer using:
1. the CURRENT PIPELINE STATE as the primary truth
2. retrieved governance evidence as supporting context
3. concise, technical, demo-safe language

STRICT RULES:
- Do not invent resources, actions, metrics, or attack paths not present in the state.
- If the answer is not available in the current state, say that clearly.
- Prefer exact node IDs, action names, counts, and metrics from the state.
- Keep the answer short but meaningful.
- When useful, mention why the answer matters operationally.
- If a resource is mentioned in the question and exists in state, refer to it directly.
- Do not claim automatic execution happened unless the state proves it.

USER QUESTION:
{query}

CURRENT PIPELINE STATE:
{state_context}

RETRIEVED EVIDENCE:
{docs_text}

OUTPUT FORMAT:
ANSWER:
""".strip()

        output = call_llm(prompt)
        if not output or len(output.strip()) < 10:
            raise ValueError("LLM returned weak copilot output")

        cleaned = output.strip()
        if cleaned.upper().startswith("ANSWER:"):
            cleaned = cleaned.split(":", 1)[1].strip()

        sources = []
        for meta in metas:
            source = meta.get("source")
            category = meta.get("category")
            if source and category:
                sources.append(f"{category}/{source}")

        return CopilotResponse(
            answer=cleaned,
            grounded=len(docs) > 0,
            sources=sources,
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@app.post("/infra", response_model=InfraResponse)
def generate_infra(data: InfraRequest):
    try:
        prompt = build_gitops_prompt(data.model_dump())
        output = call_llm(prompt)

        if not output or len(output.strip()) < 10:
            raise ValueError("LLM returned weak infra output")

        return InfraResponse(
            code=output.strip(),
            format=data.format,
            title=f"Remediation for {data.node_id}",
            summary=f"{data.action} generated for {data.node_id}",
            grounded=True,
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))