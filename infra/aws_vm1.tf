
# Terraform Infra
# Node: aws_vm1
# Time: 1774587580

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

    script = <<-EOT
      sudo yum update -y
      sudo yum install -y yum-plugin-secure-path
      sudo yum-config-manager --enable AmazonLinux-V2-2022.03
      sudo yum update -y
      sudo yum install -y amazon-linux-extras
      sudo amazon-linux-extras enable kernel-latest
      sudo yum update -y
      sudo yum install -y kernel
      sudo reboot
    EOT
  }
}
