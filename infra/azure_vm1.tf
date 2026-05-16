
# Terraform Infra
# Node: azure_vm1
# Time: 1774800999

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm resize --name azure_vm1 --size Standard_DS2_v2 --resource-group example_group
    EOF
  }
}
```
