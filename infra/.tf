
# Terraform Infra
# Node: 
# Time: 1775975629

```terraform
resource "null_resource" "node_downsize" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      terraform import aws_instance.NODE_ID NODE_ID
      terraform taint aws_instance.NODE_ID
      terraform apply -auto-approve
    EOF
  }
}
```
