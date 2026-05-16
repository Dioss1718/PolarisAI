
# Terraform Infra
# Node: 
# Time: 1775971114

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'SECURE_PATCH: Node ${self.triggers.node_id} removed due to policy approval'"
  }

  provisioner "local-exec" {
    when    = create
    command = "echo 'SECURE_PATCH: Node ${self.triggers.node_id} added due to policy approval'"
  }
}
```
