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