
# Terraform Infra
# Node: azure_vm1
# Time: 1774804921

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --resource-group <resource_group_name> --name azure_vm1 --size Medium
    EOF
  }
}
```
