
# Terraform Infra
# Node: aws_vm1
# Time: 1774609432

resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      host        = "aws_vm1"
      user        = "your_username"
      private_key = file("~/.ssh/your_private_key")
    }

    script = <<EOF
      sudo apt update && sudo apt install -y unattended-upgrades
      sudo unattended-upgrades -y --auto-upgrade
      sudo apt update && sudo apt install -y aws-secure-locations
      sudo aws secure-locations update --location-type "aws-vm1"
    EOF
  }
}
