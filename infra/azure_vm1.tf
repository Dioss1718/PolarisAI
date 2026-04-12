
# Terraform Infra
# Node: azure_vm1
# Time: 1774592038

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

resource "azurerm_resource_group" "example" {
  name     = "example-resource-group"
  location = "West US"
}

resource "azurerm_virtual_machine" "example" {
  name                  = "azure_vm1"
  resource_group_name  = azurerm_resource_group.example.name
  location              = azurerm_resource_group.example.location
  vm_size               = "Standard_DS2_v2"
}
```
