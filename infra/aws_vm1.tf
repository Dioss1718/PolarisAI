
# Terraform Infra
# Node: aws_vm1
# Time: 1774805568

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
      sudo yum install -y yum-utils
      sudo yum-config-manager --enable rhui-REGION-rhel-server-extras
      sudo yum install -y kernel-uek
      sudo yum install -y kernel-uek-devel
    EOT
  }
}
