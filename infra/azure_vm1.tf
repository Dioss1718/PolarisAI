
# Terraform Infra
# Node: azure_vm1
# Time: 1774596426

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
    action   = "DOWNSIZE_MEDIUM"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --resource-group ${azurerm_resource_group.example.name} --name azure_vm1 --size Standard_DS2_v2
    EOF
  }
}
```
