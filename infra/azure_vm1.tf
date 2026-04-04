
# Terraform Infra
# Node: azure_vm1
# Time: 1774805211

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
az vm resize --name azure_vm1 --resource-group default --size Medium
EOF
  }
}
```
