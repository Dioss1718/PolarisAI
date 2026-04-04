
# Terraform Infra
# Node: azure_vm1
# Time: 1774805557

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOT
      az vm resize --name azure_vm1 --resource-group default-rg --size Standard_DS2_v2
    EOT
  }
}
```
