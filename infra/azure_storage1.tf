
# Terraform Infra
# Node: azure_storage1
# Time: 1775967290

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Node=azure_storage1 | Action=TERMINATE_FORCE | Score=0.76 | Insight=Grounded in sla/gcp_cloud_storage.txt"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm stop --ids $(az vm list --query "[?name=='azure_storage1'].id" --output tsv)
      az vm delete --ids $(az vm list --query "[?name=='azure_storage1'].id" --output tsv)
    EOF
  }
}
```
