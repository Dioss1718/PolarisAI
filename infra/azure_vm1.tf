
# Terraform Infra
# Node: azure_vm1
# Time: 1774801094

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm resize --resource-group <resource_group> --name azure_vm1 --size Standard_DS2_v2
    EOF
  }
}
```
