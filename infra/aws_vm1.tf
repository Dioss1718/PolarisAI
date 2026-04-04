
# Terraform Infra
# Node: aws_vm1
# Time: 1774808390

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

    inline = [
      "sudo yum update -y",
      "sudo yum install -y yum-utils",
      "sudo yum-config-manager --enable rhui-REGION-rhel-server-extras",
      "sudo yum install -y kernel-$(uname -r)",
      "sudo yum install -y kernel-tools-$(uname -r)",
      "sudo yum install -y kernel-core-$(uname -r)",
      "sudo yum install -y kernel-modules-$(uname -r)",
      "sudo yum install -y kernel-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-debug-$(uname -r)",
      "sudo yum install -y kernel-debug-core-$(uname -r)",
      "sudo yum install -y kernel-debug-modules-$(uname -r)",
      "sudo yum install -y kernel-debug-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-debuginfo-$(uname -r)",
      "sudo yum install -y kernel-debuginfo-core-$(uname -r)",
      "sudo yum install -y kernel-debuginfo-modules-$(uname -r)",
      "sudo yum install -y kernel-debuginfo-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-devel-$(uname -r)",
      "sudo yum install -y kernel-devel-core-$(uname -r)",
      "sudo yum install -y kernel-devel-modules-$(uname -r)",
      "sudo yum install -y kernel-devel-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-headers-$(uname -r)",
      "sudo yum install -y kernel-headers-core-$(uname -r)",
      "sudo yum install -y kernel-headers-modules-$(uname -r)",
      "sudo yum install -y kernel-headers-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-core-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-modules-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-core-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-modules-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-core-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-modules-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-core-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-modules-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-modules-extra-$(uname -r)",
      "sudo yum install -y kernel-tools-libs-devel-headers-$(uname -r)",
      "sudo yum install -y kernel-tools
