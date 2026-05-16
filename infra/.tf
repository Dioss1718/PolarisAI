
# Terraform Infra
# Node: 
# Time: 1775972267

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'SECURE_RESTRICT: Node removed due to policy approved'"
  }

  provisioner "local-exec" {
    when    = create
    command = "echo 'SECURE_RESTRICT: Node added due to policy approved'"
  }
}
```
