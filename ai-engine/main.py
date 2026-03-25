from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from retriever import retrieve
from prompt import build_prompt
from llm import call_llm

app = FastAPI()


# 🔥 STRONG CONTRACT (CRITICAL)
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
    sources: list[str]


@app.post("/explain")
def explain(data: dict):
    try:
        query = f"{data['action']} on {data['node_type']}"

        docs = retrieve(query, data["node_type"])

        prompt = build_prompt(data, docs)

        output = call_llm(prompt)

        if not output or len(output.strip()) < 20:
            raise ValueError("LLM returned empty output")

        return {"explanation": output}

    except Exception as e:
        print("🔥 ERROR:", str(e))
        raise HTTPException(status_code=500, detail=str(e))