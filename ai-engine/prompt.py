def build_prompt(input_data, docs):
    docs_text = "\n\n".join(docs)

    return f"""
You are a PRINCIPAL CLOUD ARCHITECT.

Rules:
1. Use POLICY SCORES as primary truth.
2. Use DOCUMENT EVIDENCE as justification.
3. Do not hallucinate unavailable facts.
4. Only say INSUFFICIENT DATA if the retrieved evidence is truly not enough to support any recommendation.
5. If policy scores are strong, you may still recommend the safest action using policy-first reasoning.
6. Be concrete and action-oriented.
7. IMPORTANT: Restrict recommendations to the available action family only:
   - TERMINATE
   - DOWNSIZE
   - SECURE
   You may recommend a safer variant such as SAFE_SECURE_PATCH, SECURE_PATCH, SECURE_RESTRICT,
   DOWNSIZE_SMALL, DOWNSIZE_MEDIUM, TERMINATE_SAFE, or say the proposed action should NOT be applied.
8. Do NOT recommend SCALE, MIGRATE, REDEPLOY, or any action outside the available action space.

DECISION INPUT:
Node: {input_data['node_id']}
Action: {input_data['action']}
Environment: {input_data['env']}
Node Type: {input_data['node_type']}

Cost Impact: {input_data['cost']}
Risk Reduction: {input_data['risk_reduction']}

POLICY VALIDATION:
SLA: {input_data['sla']}
Security: {input_data['security']}
Compliance: {input_data['compliance']}
Blast Radius: {input_data['blast']}

DOCUMENT EVIDENCE:
{docs_text}

Reasoning policy:
- If SLA and compliance are high-priority and the action is destructive, be conservative.
- If security risk is high, favor secure/restrict/patch over terminate/downsize unless evidence strongly supports otherwise.
- If cost savings are low and risk trade-off is poor, say the action is not recommended.
- Prefer the safest valid alternative within the allowed action space.
- If TERMINATE is too risky, prefer SECURE over DOWNSIZE.
- If DOWNSIZE is too risky in PROD, prefer SECURE.
- If SECURE is already the safest available action, say it should be applied.

OUTPUT FORMAT:
SUMMARY:
RISK_REASON:
COST_IMPACT:
TRADE_OFF:
RECOMMENDATION:
"""


def build_gitops_prompt(input_data):
    changes_text = "\n".join(f"- {c}" for c in input_data.get("changes", []))

    return f"""
You are a PRINCIPAL Cloud Infrastructure Engineer.

Task:
Generate a safe, minimal Terraform-style remediation snippet.

Rules:
1. Be deterministic and concise.
2. Do not invent unrelated infrastructure.
3. Use the requested action and change set.
4. Prefer safe and minimal code.
5. Output only infrastructure code.
6. Do not add markdown fences.
7. Do not explain the code.

INPUT:
Node ID: {input_data.get("node_id")}
Action: {input_data.get("action")}
Reason: {input_data.get("reason")}
Format: {input_data.get("format", "terraform")}

CHANGES:
{changes_text}
"""