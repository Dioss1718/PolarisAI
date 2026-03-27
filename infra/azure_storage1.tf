
# Terraform Infra
# Node: azure_storage1
# Time: 1774600887

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az storage container delete --name azure_storage1 --yes
      az storage account delete --name azure_storage1 --yes
EOF
  }
}
```
