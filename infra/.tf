
# Terraform Infra
# Node: 
# Time: 1775975607

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm patch --ids ${node_id} --mode Secure
    EOF
  }
}
```
