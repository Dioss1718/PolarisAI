
# Terraform Infra
# Node: aws_vm1
# Time: 1774440793


resource "null_resource" "aws_vm1" {
  provisioner "local-exec" {
    command = "echo default infra applied"
  }
}

