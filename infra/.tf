
# Terraform Infra
# Node: 
# Time: 1775971127

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = "destroy"
    command = "echo 'Node ${self.triggers.node_id} has been removed.'"
  }

  provisioner "local-exec" {
    when    = "create"
    command = "echo 'Node ${self.triggers.node_id} has been created.'"
  }
}
```
