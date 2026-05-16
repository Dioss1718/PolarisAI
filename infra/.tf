
# Terraform Infra
# Node: 
# Time: 1775972253

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Node removed'"
  }

  provisioner "local-exec" {
    when    = create
    command = "echo 'Node added'"
  }
}
```
