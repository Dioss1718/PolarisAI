
# Terraform Infra
# Node: 
# Time: 1775975600

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'SECURE_PATCH: Node removed'"
  }

  provisioner "local-exec" {
    when    = create
    command = "echo 'SECURE_PATCH: Node added'"
  }
}
```
