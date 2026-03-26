def build_prompt(input_data, docs):
    docs_text = "\n\n".join(docs)

    return f"""
You are a PRINCIPAL CLOUD ARCHITECT.

Strict rules:
1. Use POLICY SCORES as primary truth.
2. Use DOCUMENT EVIDENCE as justification.
3. Do not hallucinate unavailable facts.
4. If evidence is weak, say INSUFFICIENT DATA.
5. Be concise but concrete.

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

OUTPUT FORMAT:
SUMMARY:
RISK_REASON:
COST_IMPACT:
TRADE_OFF:
RECOMMENDATION:
"""