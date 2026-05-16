
# Terraform Infra
# Node: azure_vm1
# Time: 1774805680

```terraform
resource "null_resource" "azure_vm1" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOT
      az vm resize --name azure_vm1 --resource-group default-rg --size Standard_DS2_v2
    EOT
  }

  provisioner "local-exec" {
    when    = refresh
    command = <<EOT
      az vm show --name azure_vm1 --resource-group default-rg --query "properties.hardwareProfile.vmSize" --output tsv
    EOT
  }
}

resource "null_resource" "azure_vm1_cost" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOT
      az vm deallocate --name azure_vm1 --resource-group default-rg
      az vm start --name azure_vm1 --resource-group default-rg
    EOT
  }

  provisioner "local-exec" {
    when    = refresh
    command = <<EOT
      az vm show --name azure_vm1 --resource-group default-rg --query "properties.hardwareProfile.vmSize" --output tsv
    EOT
  }
}
```
