
# Terraform Infra
# Node: aws_vm1
# Time: 1774601250

resource "aws_ssm_patch_baseline" "example" {
  name              = "example-baseline"
  description       = "Example patch baseline for safe remediation"
  approved_patches  = ["PatchID1", "PatchID2"]
  approved_patches_batch = ["PatchID3"]
  rejected_patches = ["PatchID4"]
  rejected_patches_action = "REJECT"
  description = "Patch baseline for ${aws_vm1}"
}

resource "aws_ssm_association" "example" {
  name = "AWS-ConfigureAWSPackage"
  targets {
    key    = "tag:Name"
    values = ["${aws_vm1}"]
  }
}

resource "aws_ssm_send_command" "example" {
  document_name = "AWS-RunPatchBaseline"
  targets {
    key    = "tag:Name"
    values = ["${aws_vm1}"]
  }
  parameters = {
    "Operation"        = "Install"
    "PatchBaselineId"  = aws_ssm_patch_baseline.example.id
  }
}
