
# Terraform Infra
# Node: azure_vm1
# Time: 1774805006

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --name azure_vm1 --resource-group default --size Standard_DS2_v2
    EOF
  }
}
```
