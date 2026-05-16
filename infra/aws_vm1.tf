
# Terraform Infra
# Node: aws_vm1
# Time: 1775933145


resource "null_resource" "aws_vm1" {
  provisioner "local-exec" {
    command = "echo applied SAFE_SECURE_PATCH"
  }
}

