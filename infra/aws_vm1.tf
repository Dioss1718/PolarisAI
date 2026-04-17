
# Terraform Infra
# Node: aws_vm1
# Time: 1774600698

resource "aws_ssm_patch_baseline" "example" {
  name              = "example-baseline"
  description       = "Example patch baseline"
  approved_patches  = ["patch-id-1", "patch-id-2"]
  approved_patches_batch = [
    {
      title       = "Example patch batch"
      description = "Example patch batch description"
      patch_set   = ["patch-id-3", "patch-id-4"]
    }
  ]
  rejected_patches  = ["patch-id-5", "patch-id-6"]
  rejected_patches_batch = [
    {
      title       = "Example rejected patch batch"
      description = "Example rejected patch batch description"
      patch_set   = ["patch-id-7", "patch-id-8"]
    }
  ]
}

resource "aws_ssm_association" "example" {
  name = "AWS-ConfigureAWSPackage"
  targets {
    key    = "tag:aws:ec2:instance-id"
    values = ["${aws_vm1.id}"]
  }
  parameters = {
    "Action" = "SAFE_SECURE_PATCH"
    "Reason" = "Policy Approved"
  }
}
