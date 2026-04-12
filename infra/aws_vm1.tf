
# Terraform Infra
# Node: aws_vm1
# Time: 1774801536

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

    script = <<-EOT
      sudo apt-get update
      sudo apt-get install -y unattended-upgrades
      sudo unattended-upgrades -y --auto-all
      sudo apt-get -y update
      sudo apt-get -y upgrade
      sudo apt-get -y dist-upgrade
      sudo apt-get -y autoremove
      sudo apt-get -y autoclean
      sudo apt-get -y clean
      sudo apt-get -y autoremove --purge
      sudo apt-get -y autoclean --purge
      sudo apt-get -y clean --purge
      sudo apt-get -y autoremove --purge && sudo apt-get -y autoclean --purge && sudo apt-get -y clean --purge
    EOT
  }
}
