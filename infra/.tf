
# Terraform Infra
# Node: 
# Time: 1775971110

```terraform
resource "null_resource" "example" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    command = <<EOF
      aws ec2 modify-instance-attribute --instance-id NODE_ID --attribute instanceSecurityGroups.GroupId --value "sg-12345678"
    EOF
  }
}
```
