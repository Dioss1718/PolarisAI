
# Terraform Infra
# Node: aws_vm1
# Time: 1774817733

resource "aws_ssm_patch_baseline" "vm1_baseline" {
  name              = "vm1_baseline"
  description       = "Baseline for vm1"
  approval_rule {
    approval_type = "Manual"
  }
}

resource "aws_ssm_association" "vm1_association" {
  name               = aws_ssm_patch_baseline.vm1_baseline.name
  targets {
    key    = "tag:NodeID"
    values = ["aws_vm1"]
  }
}
