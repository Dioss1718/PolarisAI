
# Terraform Infra
# Node: aws_vm1
# Time: 1774600898

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
      "sudo apt update",
      "sudo apt install -y unattended-upgrades",
      "sudo unattended-upgrades -y --download-only",
      "sudo apt dist-upgrade -y",
      "sudo apt autoremove -y",
      "sudo apt autoclean -y",
      "sudo apt full-upgrade -y",
      "sudo shutdown -r now"
    ]
  }
}
