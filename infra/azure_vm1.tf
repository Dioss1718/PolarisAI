
# Terraform Infra
# Node: azure_vm1
# Time: 1775929183

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm downgrade --name azure_vm1 --size Standard_DS2_v2
    EOF
  }
}
```
