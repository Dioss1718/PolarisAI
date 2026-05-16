
# Terraform Infra
# Node: azure_storage1
# Time: 1774818330

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id   = "azure_storage1"
    action    = "TERMINATE_FORCE"
    reason    = "Policy Approved"
    exposure  = "PRIVATE"
    utilization = "0.00"
    cost       = "0.00"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Remediation applied for azure_storage1'"
  }
}
```
