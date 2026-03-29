
# Terraform Infra
# Node: azure_vm1
# Time: 1774810713

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --name azure_vm1 --resource-group default --size Medium
    EOF
  }
}
```
