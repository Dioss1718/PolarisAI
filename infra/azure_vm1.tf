
# Terraform Infra
# Node: azure_vm1
# Time: 1774456713


resource "null_resource" "azure_vm1" {
  provisioner "local-exec" {
    command = "echo default infra applied"
  }
}

