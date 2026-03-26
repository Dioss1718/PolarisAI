
# Terraform Infra
# Node: aws_vm1
# Time: 1774534029


resource "null_resource" "aws_vm1" {
  provisioner "local-exec" {
    command = "echo applied SAFE_SECURE_PATCH"
  }
}

