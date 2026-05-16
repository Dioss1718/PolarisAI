
# Terraform Infra
# Node: 
# Time: 1775971074

```terraform
resource "null_resource" "remediation" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      terraform taint aws_instance.NODE_ID
      terraform apply -auto-approve
      terraform taint aws_instance.NODE_ID
      terraform apply -auto-approve
    EOF
  }

  provisioner "local-exec" {
    when    = refresh
    command = <<EOF
      terraform taint aws_instance.NODE_ID
      terraform apply -auto-approve
    EOF
  }

  connection {
    command = "sleep 1"
  }
}

resource "aws_instance" "NODE_ID" {
  // existing instance configuration
}

resource "aws_instance" "NODE_ID" {
  // new instance configuration
  instance_type = "t2.micro"
}
```
