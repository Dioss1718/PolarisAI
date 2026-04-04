
# Terraform Infra
# Node: aws_vm1
# Time: 1774599269

resource "aws_ssm_patch_baseline" "vm1_baseline" {
  name              = "vm1_baseline"
  description       = "Baseline for vm1"
  approved_patches  = ["patch-1234567890"]
  rejected_patches  = ["patch-1234567891"]
  approved_patches_calendar = ["2022-01"]
  description       = "Baseline for vm1"
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
  content         = jsonencode({
    "schemaVersion": "1.2",
    "description": "Patch vm1",
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
  })
}

resource "aws_ssm_patch_group" "vm1_group" {
  name = "vm1_group"
}

resource "aws_ssm_maintenance_window" "vm1_window" {
  name          = "vm1_window"
  description   = "Maintenance window for vm1"
  schedule      = "cron(0 12 * * ? *)"
  duration      = 1
  start_time    = "2022-01-01T12:00:00Z"
  end_time      = "2022-01-01T13:00:00Z"
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
