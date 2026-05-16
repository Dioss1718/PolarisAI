
# Terraform Infra
# Node: aws_vm1
# Time: 1774813968

resource "aws_ssm_association" "vm1" {
  name          = "AWS-SecurePatch"
  targets {
    key    = "tag:Name"
    values = ["aws_vm1"]
  }
  parameters = {
    "PatchGroup" = "Default"
  }
}
