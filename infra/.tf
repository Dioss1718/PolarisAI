
# Terraform Infra
# Node: 
# Time: 1775975626

```terraform
resource "null_resource" "node_downsize" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    command = <<EOF
      terraform apply -auto-approve -var "node_id=NODE_ID" -var "action=DOWNSIZE_MEDIUM" -var "reason=Policy Approved"
    EOF
  }
}
```
