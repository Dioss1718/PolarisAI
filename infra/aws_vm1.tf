
# Terraform Infra
# Node: aws_vm1
# Time: 1774590153

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
      "sudo unattended-upgrades -y --auto-upgrade"
    ]
  }
}
