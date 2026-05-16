
# Terraform Infra
# Node: aws_vm1
# Time: 1775929188

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
      "sudo apt update && sudo apt install -y unattended-upgrades",
      "sudo unattended-upgrades -y --auto-upgrade",
      "sudo apt update && sudo apt install -y aws-cfn-bootstrap",
      "sudo aws-cfn-bootstrap install-cfn-hup",
      "sudo cfn-hup -e /etc/cfn/cfn-hup.conf"
    ]
  }
}
