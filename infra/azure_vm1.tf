
# Terraform Infra
# Node: azure_vm1
# Time: 1774832506

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --resource-group ${azurerm_resource_group.example.name} --name ${azurerm_linux_virtual_machine.azure_vm1.name} --size Standard_DS2_v2
    EOF
  }
}
```

```terraform
resource "azurerm_resource_group" "example" {
  name     = "example-resource-group"
  location = "West US"
}

resource "azurerm_linux_virtual_machine" "azure_vm1" {
  name                = "example-vm"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  size                = "Standard_DS2_v2"
}
```
