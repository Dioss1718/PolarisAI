
# Terraform Infra
# Node: 
# Time: 1775971083

```terraform
resource "null_resource" "node_termination" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Terminating Node: NODE_ID due to Policy Approved'"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "aws ec2 terminate-instances --instance-ids NODE_ID"
  }
}
```
