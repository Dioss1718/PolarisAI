
# Terraform Infra
# Node: azure_storage1
# Time: 1774805637

resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az storage container delete --name azure_storage1 --yes
    EOF
  }
}
