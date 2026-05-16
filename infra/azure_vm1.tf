
# Terraform Infra
# Node: azure_vm1
# Time: 1774802043

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --ids ${azurerm_virtual_machine.azure_vm1.id} --size "Standard_DS2_v2"
    EOF
  }
}
```

```terraform
resource "null_resource" "azure_vm1_update_tags" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm update --ids ${azurerm_virtual_machine.azure_vm1.id} --set "tags.cost=136.50, tags.utilization=10.50"
    EOF
  }
}
```
