
# Terraform Infra
# Node: azure_vm1
# Time: 1774809238

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    reason  = "Policy Approved"
    action  = "DOWNSIZE_MEDIUM"
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --resource-group ${azurerm_resource_group.example.name} --name azure_vm1 --size Standard_DS2_v2
    EOF
  }
}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West US"
}
```
