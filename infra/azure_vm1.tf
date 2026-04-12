
# Terraform Infra
# Node: azure_vm1
# Time: 1774592360

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --ids ${azurerm_virtual_machine.azure_vm1.id} --size Standard_DS2_v2
    EOF
  }
}
```
