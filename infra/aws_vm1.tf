
# Terraform Infra
# Node: aws_vm1
# Time: 1774818337

resource "aws_ssm_patch_baseline" "vm1_baseline" {
  name              = "vm1_baseline"
  description       = "vm1 patch baseline"
  approved_patches  = ["patch-1234567890"]
  rejected_patches  = ["patch-9876543210"]
  approved_patches_batch = ["patch-1234567890"]
  patch_set = "APPLICABLE"
}

resource "aws_ssm_association" "vm1_association" {
  name = "AWS-ConfigureAWSPackage"
  targets {
    key    = "tag:aws:ec2:instance-id"
    values = ["${aws_instance.vm1.id}"]
  }
}

resource "aws_ssm_document" "vm1_document" {
  name            = "vm1_document"
  document_type   = "Command"
  content         = <<EOF
  {
    "schemaVersion": "1.2",
    "description": "vm1 document",
    "mainSteps": [
      {
        "action": "aws:runShellScript",
        "name": "runShellScript",
        "inputs": {
          "runCommand": [
            "sudo yum update -y",
            "sudo yum install -y aws-cfn-bootstrap"
          ]
        }
      }
    ]
  }
  EOF
}

resource "aws_ssm_patch_group" "vm1_patch_group" {
  name = "vm1_patch_group"
}

resource "aws_ssm_maintenance_window" "vm1_window" {
  name          = "vm1_window"
  description   = "vm1 maintenance window"
  schedule      = "cron(0 1 * * ? *)"
  start_time    = "2024-03-16T00:00:00Z"
  end_time      = "2024-03-16T02:00:00Z"
  allow_unassociated_targets = true
}

resource "aws_ssm_maintenance_window_target" "vm1_target" {
  window_id = aws_ssm_maintenance_window.vm1_window.id
  targets {
    key    = "tag:aws:ec2:instance-id"
    values = ["${aws_instance.vm1.id}"]
  }
}

resource "aws_ssm_maintenance_window_task" "vm1_task" {
  window_id = aws_ssm_maintenance_window.vm1_window.id
  task_type = "RUN_COMMAND"
  task_arn  = aws_ssm_document.vm1_document.arn
}
