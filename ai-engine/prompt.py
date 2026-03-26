def build_prompt(input_data, docs):
    docs_text = "\n\n".join(docs) if docs else "NO DOCUMENT EVIDENCE RETRIEVED."

    return f"""
You are a PRINCIPAL CLOUD ARCHITECT.

Your job is to explain and validate an infrastructure action using:
1. policy scores as primary truth
2. retrieved document evidence as support
3. strict action-space constraints

CRITICAL RULES:
1. Do NOT hallucinate unavailable facts.
2. Do NOT recommend actions outside the allowed action space.
3. Allowed action families are ONLY:
   - TERMINATE
   - DOWNSIZE
   - SECURE
4. Allowed action variants include:
   - TERMINATE_SAFE
   - TERMINATE_FORCE
   - DOWNSIZE_SMALL
   - DOWNSIZE_MEDIUM
   - SECURE_PATCH
   - SECURE_RESTRICT
   - SAFE_SECURE_PATCH
5. Do NOT recommend SCALE, MIGRATE, REDEPLOY, RESTART, or any action outside the available action space.
6. If the proposed action is weak or risky, recommend the safest valid alternative ONLY within the allowed action space.
7. Only say "INSUFFICIENT DATA" when both:
   - the retrieved evidence is truly missing or irrelevant, AND
   - policy scores are too weak or conflicting to support even a conservative recommendation.
8. If policy scores are strong, you MUST still provide a best recommendation from the allowed action space.

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

REASONING POLICY:
- If the environment is PROD and the action is destructive, be conservative.
- If SLA or compliance importance is high, avoid destructive recommendations unless strongly justified.
- If security risk is high, prefer SECURE over DOWNSIZE or TERMINATE.
- If DOWNSIZE has weak benefit and meaningful operational risk, recommend SECURE.
- If TERMINATE is risky for availability, recommend SECURE.
- If SECURE is already the safest available action, recommend SECURE.
- If no retrieved evidence is useful but policy is strong, use policy-first reasoning and still recommend the safest allowed action.
- NEVER recommend SCALE.

OUTPUT FORMAT:
SUMMARY:
RISK_REASON:
COST_IMPACT:
TRADE_OFF:
RECOMMENDATION:
"""