
# Terraform Infra
# Node: aws_vm1
# Time: 1774813128

resource "null_resource" "aws_vm1" {
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
      "sudo apt-get update",
      "sudo apt-get install -y unattended-upgrades",
      "sudo apt-get install -y awscli",
      "sudo aws ssm send-command --instance-id aws_vm1 --document-name AWS-UpdateSSMDocument --parameters '{"Command": "update-ssm-document"}'",
      "sudo aws ssm send-command --instance-id aws_vm1 --document-name AWS-UpdateSSMDocument --parameters '{"Command": "update-ssm-document"}'",
      "sudo aws ssm send-command --instance-id aws_vm1 --document-name AWS-UpdateSSMDocument --parameters '{"Command": "update-ssm-document"}'",
    ]
  }
}
