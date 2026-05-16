
# Terraform Infra
# Node: azure_vm1
# Time: 1774808768

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm resize --resource-group ${azurerm_resource_group.example.name} --name azure_vm1 --size Standard_DS2_v2
    EOF
  }
}
```
