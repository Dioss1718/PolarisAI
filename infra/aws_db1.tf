# Terraform Infra
# Node: aws_db1
# Time: 1776077958

resource "null_resource" "aws_db1_SECURE_PATCH" {
  triggers = {
    reason = "SECURE_PATCH on aws_db1 | RiskReduction=2.61 | CostImpact=231.00 | Penalty=0.33 | FinalScore=-0.09 | Node=aws_db1 | Action=SECURE_PATCH | PolicyScore=0.89 | DimensionScores[SLA=1.00, Security=0.73, Compliance=1.00, Blast=0.80] | Insight=Grounded in sla/azure_global.txt | AppliedRules=security.restrict-or-patch-boost: Security-hardening action improves posture."
  }
}
