
# Terraform Infra
# Node: 
# Time: 1775971118

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      aws ec2 modify-instance-attribute --instance-id ${self.triggers.node_id} --attribute instanceType --value "t3.micro"
      aws ec2 stop-instances --instance-ids ${self.triggers.node_id}
      aws ec2 start-instances --instance-ids ${self.triggers.node_id}
    EOF
  }
}
```
