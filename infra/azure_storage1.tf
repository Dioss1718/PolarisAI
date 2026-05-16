
# Terraform Infra
# Node: azure_storage1
# Time: 1774810710

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    reason  = "Policy Approved"
    action  = "TERMINATE_FORCE"
    node_id = "azure_storage1"
  }

  provisioner "local-exec" {
    command = <<EOF
      az storage container delete --name azure_storage1
    EOF
  }
}
```
