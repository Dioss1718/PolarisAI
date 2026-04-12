
# Terraform Infra
# Node: aws_vm1
# Time: 1774805646

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
      sudo yum update -y
      sudo yum install -y yum-utils
      sudo yum-config-manager --enable rhui-REGION-rhel-server-extras
      sudo yum install -y kernel-uek
      sudo yum install -y kernel-uek-devel
      sudo yum install -y kernel-uek-modules
      sudo yum install -y kernel-uek-modules-extra
      sudo yum install -y kernel-uek-firmware
      sudo yum install -y kernel-uek-firmware-modules
      sudo yum install -y kernel-uek-firmware-modules-extra
      sudo yum install -y kernel-uek-modules-$(uname -r)
      sudo yum install -y kernel-uek-modules-extra-$(uname -r)
      sudo yum install -y kernel-uek-firmware-$(uname -r)
      sudo yum install -y kernel-uek-firmware-modules-$(uname -r)
      sudo yum install -y kernel-uek-firmware-modules-extra-$(uname -r)
      sudo yum install -y kernel-uek-modules-$(uname -r)-debug
      sudo yum install -y kernel-uek-modules-extra-$(uname -r)-debug
      sudo yum install -y kernel-uek-firmware-$(uname -r)-debug
      sudo yum install -y kernel-uek-firmware-modules-$(uname -r)-debug
      sudo yum install -y kernel-uek-firmware-modules-extra-$(uname -r)-debug
      sudo yum install -y kernel-uek-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-modules-extra-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-$(uname -r)-debuginfo
      sudo yum install -y kernel-uek-firmware-modules-$(uname -r
