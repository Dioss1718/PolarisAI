
# Terraform Infra
# Node: aws_vm1
# Time: 1774546407

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
      "sudo apt update",
      "sudo apt install -y unattended-upgrades",
      "sudo apt install -y aws-secure-locator",
      "sudo aws-secure-locator --install",
      "sudo aws-secure-locator --enable",
      "sudo apt update",
      "sudo apt install -y aws-secure-locator-updater",
      "sudo aws-secure-locator-updater --install",
      "sudo aws-secure-locator-updater --enable",
    ]
  }
}
