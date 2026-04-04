
# Terraform Infra
# Node: azure_storage1
# Time: 1774768157

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm stop --ids $(az vm show --name azure_storage1 --resource-group <RG_NAME> --query id --output tsv)
      az vm delete --ids $(az vm show --name azure_storage1 --resource-group <RG_NAME> --query id --output tsv)
    EOF
  }
}
```
