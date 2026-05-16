
# Terraform Infra
# Node: aws_vm1
# Time: 1775964708


resource "null_resource" "aws_vm1" {
  provisioner "local-exec" {
    command = "echo applied SECURE_PATCH"
  }
}

