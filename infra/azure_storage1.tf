
# Terraform Infra
# Node: azure_storage1
# Time: 1774808375


resource "null_resource" "azure_storage1" {
  provisioner "local-exec" {
    command = "echo applied TERMINATE_FORCE"
  }
}

