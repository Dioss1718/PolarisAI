
# Terraform Infra
# Node: 
# Time: 1775972278

```terraform
resource "null_resource" "node_downsize" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    command = <<EOF
      terraform taint -node "NODE_ID"
      terraform apply -auto-approve
    EOF
  }
}
```
