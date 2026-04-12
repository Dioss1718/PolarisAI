
# Terraform Infra
# Node: aws_vm1
# Time: 1775960911


resource "null_resource" "aws_vm1" {
  provisioner "local-exec" {
    command = "echo applied SECURE_PATCH"
  }
}

