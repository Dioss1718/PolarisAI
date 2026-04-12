
# Terraform Infra
# Node: 
# Time: 1775975657

```terraform
resource "null_resource" "remediation" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Remediation for Node ${self.triggers.node_id} downsized to medium due to policy approval'"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Changes: Node missing in old graph, Node missing in new graph'"
  }
}
```
