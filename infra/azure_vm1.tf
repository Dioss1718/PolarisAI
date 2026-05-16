
# Terraform Infra
# Node: azure_vm1
# Time: 1775966262


resource "null_resource" "azure_vm1" {
  provisioner "local-exec" {
    command = "echo applied DOWNSIZE_MEDIUM"
  }
}

