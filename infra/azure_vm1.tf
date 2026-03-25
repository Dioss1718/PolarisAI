
# Terraform Infra
# Node: azure_vm1
# Time: 1774440783


resource "null_resource" "azure_vm1" {
  provisioner "local-exec" {
    command = "echo default infra applied"
  }
}

