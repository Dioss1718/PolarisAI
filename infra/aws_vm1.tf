
# Terraform Infra
# Node: aws_vm1
# Time: 1774802148

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
      "sudo aws ssm send-command --instance-id aws_vm1 --command 'sudo apt-get update && sudo apt-get install -y unattended-upgrades && sudo apt-get install -y awscli'",
      "sudo aws ssm send-command --instance-id aws_vm1 --command 'sudo yum update -y'",
      "sudo aws ssm send-command --instance-id aws_vm1 --command 'sudo yum install -y epel-release'",
      "sudo aws ssm send-command --instance-id aws_vm1 --command 'sudo yum install -y awscli'",
    ]
  }
}
