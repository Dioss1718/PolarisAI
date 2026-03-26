from typing import List, Optional

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


@app.get("/health")
def health():
    return {"status": "ok"}


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