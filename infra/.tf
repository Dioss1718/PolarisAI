
# Terraform Infra
# Node: 
# Time: 1775975633

```terraform
resource "null_resource" "node_remediation" {
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

  provisioner "local-exec" {
    when    = create
    command = <<EOF
      terraform import aws_instance.NODE_ID NODE_ID
      terraform apply -auto-approve
    EOF
  }
}
```
