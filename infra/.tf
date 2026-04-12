
# Terraform Infra
# Node: 
# Time: 1775975636

```terraform
resource "null_resource" "node_termination" {
  triggers = {
    node_id = "<NODE_ID>"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'TERMINATE_SAFE: ${self.triggers.node_id} - Reason: Policy Approved' >> /tmp/terraform_destroy_log.txt"
  }

  connection {
    # Add connection details as needed
  }

  lifecycle {
    create_before_destroy = true
  }
}
```
