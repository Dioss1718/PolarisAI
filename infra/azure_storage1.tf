
# Terraform Infra
# Node: azure_storage1
# Time: 1774587573

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    action = "TERMINATE_FORCE"
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az storage container delete --name azure_storage1 --yes
EOF
  }
}
```
