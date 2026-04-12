
# Terraform Infra
# Node: 
# Time: 1775975646

```terraform
resource "null_resource" "remediation" {
  triggers = {
    node_id = "NODE_ID"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      terraform taint aws_instance.NODE_ID
      terraform apply -auto-approve
    EOF
  }
}
```
