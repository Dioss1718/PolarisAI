
# Terraform Infra
# Node: azure_vm1
# Time: 1774817721

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

  connection {
    type        = "ssh"
    host        = azurerm_public_ip.example.ip_address
    user        = "username"
    private_key = file("~/.ssh/azure_key")
  }

  depends_on = [
    azurerm_resource_group.example,
    azurerm_public_ip.example,
    azurerm_key_vault_secret.example,
  ]
}

resource "azurerm_resource_group" "example" {
  name     = "example-resource-group"
  location = "West US"
}

resource "azurerm_public_ip" "example" {
  name                = "example-public-ip"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  allocation_method   = "Dynamic"
}

resource "azurerm_key_vault_secret" "example" {
  name         = "example-secret"
  value        = "example-secret-value"
  key_vault_id = azurerm_key_vault.example.id
}

resource "azurerm_key_vault" "example" {
  name                = "example-key-vault"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  sku_name            = "standard"
}
```
