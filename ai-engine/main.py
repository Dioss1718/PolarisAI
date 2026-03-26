from typing import List
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

from retriever import retrieve
from prompt import build_prompt
from llm import call_llm

app = FastAPI()


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
            f"{data.action} for {data.node_type} in {data.env} "
            f"cost impact security implications SLA compliance"
        )

        docs, metas = retrieve(query, data.node_type, data.action)
        prompt = build_prompt(data.model_dump(), docs)
        output = call_llm(prompt)

        if not output or len(output.strip()) < 20:
            raise ValueError("LLM returned weak output")

        sources = []
        for meta in metas:
            source = meta.get("source")
            category = meta.get("category")
            if source and category:
                sources.append(f"{category}/{source}")

        return ExplainResponse(
            explanation=output,
            grounded=len(docs) > 0,
            sources=sources
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))