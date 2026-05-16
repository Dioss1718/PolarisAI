
# Terraform Infra
# Node: aws_vm1
# Time: 1774833202

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

    script = <<-EOF
      sudo apt update && sudo apt install -y unattended-upgrades
      sudo unattended-upgrades -y --auto-upgrade
      sudo apt update && sudo apt install -y aws-secure-locations
      sudo aws secure-locations update --locations /etc/ssh/ssh_host_rsa_key
      sudo aws secure-locations update --locations /etc/ssh/ssh_host_ecdsa_key
      sudo aws secure-locations update --locations /etc/ssh/ssh_host_ed25519_key
    EOF
  }
}
