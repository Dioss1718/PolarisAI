
# Terraform Infra
# Node: aws_vm1
# Time: 1774809241

resource "null_resource" "aws_vm1_patch" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      host        = "aws_vm1"
      user        = "ubuntu"
      private_key = file("~/.ssh/aws_vm1_key")
    }

    inline = [
      "sudo apt update && sudo apt install -y unattended-upgrades",
      "sudo apt update && sudo apt install -y aws-ssm-agent",
      "sudo systemctl enable aws-ssm-agent && sudo systemctl start aws-ssm-agent",
      "sudo ssm send-command --document-name 'AWS-RunShellScript' --targets 'Key=instance-id,Values=$(hostname)' --parameters 'commands=[\"sudo apt update && sudo apt install -y python3\"]' --timeout-seconds 300"
    ]
  }
}
