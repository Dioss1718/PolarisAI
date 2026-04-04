
# Terraform Infra
# Node: aws_vm1
# Time: 1774805113

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

    script = <<-EOF
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
      sudo yum install -y kernel-uek-firmware-modules-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek-uek-uek-extra
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-uek-uek-uek-uek-uek-uek-uek-uek
      sudo yum install -y kernel-uek-firmware-modules-uek-uek-
