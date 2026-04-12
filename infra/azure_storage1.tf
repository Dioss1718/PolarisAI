
# Terraform Infra
# Node: azure_storage1
# Time: 1775965754


resource "null_resource" "azure_storage1" {
  provisioner "local-exec" {
    command = "echo applied TERMINATE_FORCE"
  }
}

