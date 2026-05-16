
# Terraform Infra
# Node: 
# Time: 1775972281

```terraform
resource "null_resource" "remediation" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    command = <<EOF
      terraform apply -auto-approve -var 'node_id=NODE_ID' -var 'action=DOWNSIZE_MEDIUM' -var 'reason=Policy Approved'
    EOF
  }
}
```
