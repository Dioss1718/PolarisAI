
# Terraform Infra
# Node: 
# Time: 1775975597

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }
  provisioner "local-exec" {
    command = <<EOF
      aws ec2 modify-instance-attribute --instance-id ${self.triggers.node_id} --attribute instanceSecureAndInsecureDirectMemoryAccess --value enabled
    EOF
  }
}
```
