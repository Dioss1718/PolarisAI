
# Terraform Infra
# Node: azure_storage1
# Time: 1774801214

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    action      = "TERMINATE_FORCE"
    reason      = "Policy Approved"
    node_id     = "azure_storage1"
    exposure    = "PRIVATE"
    utilization = "0.00"
    cost        = "0.00"
  }
}
```
