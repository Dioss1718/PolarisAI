
# Terraform Infra
# Node: aws_vm1
# Time: 1774440789


resource "null_resource" "aws_vm1" {
  provisioner "local-exec" {
    command = "echo default infra applied"
  }
}

