
# Terraform Infra
# Node: aws_vm1
# Time: 1775965768

```terraform
resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      aws ec2 modify-image-attribute --image-id <image-id> --attribute ec2ImageAction --operation-type SecureTheConsole --no-associate-public-ip-address
    EOF
  }
}
```
