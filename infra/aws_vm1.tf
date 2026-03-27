
# Terraform Infra
# Node: aws_vm1
# Time: 1774593605

resource "aws_ssm_association" "vm1" {
  name          = "AWS-SecurePatch"
  targets {
    key    = "tag:NodeID"
    values = ["aws_vm1"]
  }
  parameters = {
    "PatchBaseline" = "AWS-20220101"
  }
}
