
# Terraform Infra
# Node: azure_storage1
# Time: 1774801109

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
az vm deallocate --resource-group ${azurerm_resource_group.example.name} --name ${azurerm_linux_virtual_machine.example.name}
az vm stop --resource-group ${azurerm_resource_group.example.name} --name ${azurerm_linux_virtual_machine.example.name}
az vm delete --resource-group ${azurerm_resource_group.example.name} --name ${azurerm_linux_virtual_machine.example.name} --yes
az storage container delete --name ${azurerm_storage_container.example.name} --account-name ${azurerm_storage_account.example.name} --yes
EOF
  }

  connection {
    type        = "ssh"
    host        = azurerm_public_ip.example.ip_address
    user        = "username"
    private_key = file("~/.ssh/your_private_key")
  }
}

resource "azurerm_resource_group" "example" {
  name     = "example-resource-group"
  location = "West US"
}

resource "azurerm_storage_account" "example" {
  name                     = "examplestorageaccount"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "example" {
  name                     = "examplecontainer"
  resource_group_name      = azurerm_resource_group.example.name
  storage_account_name     = azurerm_storage_account.example.name
  container_access_type    = "private"
}

resource "azurerm_public_ip" "example" {
  name                = "example-public-ip"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  allocation_method   = "Dynamic"
}

resource "azurerm_linux_virtual_machine" "example" {
  name                  = "example-vm"
  resource_group_name   = azurerm_resource_group.example.name
  location              = azurerm_resource_group.example.location
  size                  = "Standard_DS2_v2"
  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Premium_LRS"
  }
  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
}
```
