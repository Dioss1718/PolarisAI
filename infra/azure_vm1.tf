
# Terraform Infra
# Node: azure_vm1
# Time: 1774593596

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --name azure_vm1 --resource-group ${azurerm_resource_group.example.name} --size Standard_DS2_v2
    EOF
  }
}

resource "null_resource" "azure_vm1_update_tags" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm update --name azure_vm1 --resource-group ${azurerm_resource_group.example.name} --set tags=<<EOF
        {
          "utilization": "10.50",
          "cost": "136.50"
        }
      EOF
    EOF
  }
}
```
