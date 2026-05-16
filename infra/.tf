
# Terraform Infra
# Node: 
# Time: 1775972271

```terraform
resource "null_resource" "node_termination" {
  triggers = {
    node_id = "Node ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'TERMINATE_SAFE: Node ID - Reason: Policy Approved' && echo 'TERMINATE_SAFE: Node ID - Reason: Policy Approved'"
  }

  connection {
    type        = "ssh"
    host        = "Node ID"
    user        = "username"
    private_key = file("~/.ssh/private_key")
  }

  connection {
    type        = "ssh"
    host        = "Node ID"
    user        = "username"
    private_key = file("~/.ssh/private_key")
  }

  depends_on = []
}
```
