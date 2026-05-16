
# Terraform Infra
# Node: azure_vm1
# Time: 1774587576

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm downgrade --ids $(az vm show --name azure_vm1 --query id --output tsv) --size Standard_DS2_v2
    EOF
  }
}
```
