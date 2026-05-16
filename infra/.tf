
# Terraform Infra
# Node: 
# Time: 1775972274

```terraform
resource "null_resource" "node_termination" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'TERMINATE_FORCE Reason: Policy Approved' && echo 'Node ID: NODE_ID'"
  }

  lifecycle {
    create_before_destroy = true
  }
}
```
