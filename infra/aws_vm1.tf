
# Terraform Infra
# Node: aws_vm1
# Time: 1775967298

```terraform
resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      aws ec2 modify-image-attribute --image-id ami-12345678 --attribute secure-patch --value enabled
    EOF
  }
}
```
