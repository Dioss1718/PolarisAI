
# Terraform Infra
# Node: aws_vm1
# Time: 1774801732

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
      sudo yum install -y kernel-uek-firmware
      sudo yum install -y kernel-uek-modules
      sudo yum install -y kernel-uek-modules-extra
      sudo yum install -y kernel-uek-headers
      sudo yum install -y kernel-uek-tools
      sudo yum install -y kernel-uek-tools-libs
      sudo yum install -y kernel-uek-tools-libs-devel
      sudo yum install -y kernel-uek-tools-libs-static
      sudo yum install -y kernel-uek-tools-static
      sudo yum install -y kernel-uek-tools-libs-compat
      sudo yum install -y kernel-uek-tools-libs-compat-devel
      sudo yum install -y kernel-uek-tools-libs-compat-static
      sudo yum install -y kernel-uek-tools-compat
      sudo yum install -y kernel-uek-tools-compat-devel
      sudo yum install -y kernel-uek-tools-compat-static
      sudo yum install -y kernel-uek-tools-compat-libs
      sudo yum install -y kernel-uek-tools-compat-libs-devel
      sudo yum install -y kernel-uek-tools-compat-libs-static
      sudo yum install -y kernel-uek-tools-compat-libs-compat
      sudo yum install -y kernel-uek-tools-compat-libs-compat-devel
      sudo yum install -y kernel-uek-tools-compat-libs-compat-static
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-devel
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-static
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-devel
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-static
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-devel
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-static
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-compat
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-compat-devel
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-compat-static
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-compat-libs
      sudo yum install -y kernel-uek-tools-compat-libs-compat-libs-compat-libs-compat-libs-devel
