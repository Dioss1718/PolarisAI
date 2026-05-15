
# Terraform Infra
# Node: azure_storage1
# Time: 1774808655

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm stop --ids $(az vm list --query "[?name=='${null_resource.azure_storage1.name}'].id" --output tsv)
      az vm deallocate --ids $(az vm list --query "[?name=='${null_resource.azure_storage1.name}'].id" --output tsv)
      az vm delete --ids $(az vm list --query "[?name=='${null_resource.azure_storage1.name}'].id" --output tsv) --yes
    EOF
  }
}
```
