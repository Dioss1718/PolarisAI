
# Terraform Infra
# Node: aws_vm1
# Time: 1774596437

resource "aws_ssm_patch_baseline" "example" {
  name              = "example-baseline"
  description       = "Example patch baseline"
  approval_rule {
    patch_rule {
      patch_filter {
        key    = "PRODUCT"
        values = ["WindowsServer*"]
      }
      approve_after_days = 7
    }
  }
}

resource "aws_ssm_association" "example" {
  name               = aws_ssm_patch_baseline.example.name
  targets {
    key    = "tag:NodeID"
    values = ["aws_vm1"]
  }
}
