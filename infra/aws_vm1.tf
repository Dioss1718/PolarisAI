
# Terraform Infra
# Node: aws_vm1
# Time: 1774805220

resource "aws_ssm_patch_baseline" "example" {
  name              = "example-patch-baseline"
  description       = "Patch baseline for safe and secure patching"
  approved_patches   = ["CVE-2022-0001", "CVE-2022-0002"]
  approved_patches_batch = ["CVE-2022-0003"]
  rejected_patches    = ["CVE-2022-0004"]
  rejected_patches_action = "DEPLOY"
  description            = "Patch baseline for safe and secure patching"
}

resource "aws_ssm_association" "example" {
  name             = "AWS-ConfigureAWSPackage"
  targets          = {
    key    = "tag:aws_vm1"
    values = ["true"]
  }
  parameters       = {
    "Action" = "SafeSecurePatch"
    "Reason" = "Policy Approved"
  }
}
