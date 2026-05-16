
# Terraform Infra
# Node: 
# Time: 1775975650

```terraform
resource "null_resource" "remediation" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      terraform apply -auto-approve -var 'node_id=NODE_ID' -var 'action=DOWNSIZE_AGGRESSIVE' -var 'reason=Policy Approved'
    EOF
  }
}
```
