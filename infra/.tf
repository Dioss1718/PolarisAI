
# Terraform Infra
# Node: 
# Time: 1775975604

```terraform
resource "null_resource" "node_removal" {
  triggers = {
    node_id = "Node ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Node removed successfully'"
  }

  provisioner "local-exec" {
    when    = create
    command = "echo 'Node creation skipped due to policy'"
  }
}
```
