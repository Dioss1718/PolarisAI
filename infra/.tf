
# Terraform Infra
# Node: 
# Time: 1775971088

```terraform
resource "null_resource" "node_termination" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'TERMINATE_FORCE: Reason: Policy Approved' > /dev/null"
  }

  provisioner "remote-exec" {
    when    = destroy
    connection {
      type        = "ssh"
      host        = "NODE_IP"
      user        = "NODE_USER"
      private_key = file("~/.ssh/NODE_KEY")
    }
    inline = [
      "echo 'TERMINATE_FORCE: Reason: Policy Approved' > /dev/null",
      "shutdown -h now"
    ]
  }
}
```
