
# Terraform Infra
# Node: azure_vm1
# Time: 1774600068

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
    action   = "DOWNSIZE_MEDIUM"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOT
      az vm downgrade --ids $(az vm show --name azure_vm1 --query id --output tsv) --size Standard_DS2_v2
    EOT
  }
}
```
