
# Terraform Infra
# Node: azure_storage1
# Time: 1774808921

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az vm stop --ids $(az vm list --query "[?name=='${null_resource.azure_storage1.name}'].id" --output tsv)
      az vm delete --ids $(az vm list --query "[?name=='${null_resource.azure_storage1.name}'].id" --output tsv)
    EOF
  }
}

resource "null_resource" "azure_storage1_network" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az network nic update --ids $(az network nic list --query "[?name=='${null_resource.azure_storage1_network.name}'].id" --output tsv) --network-security-group "NetworkSecurityGroup"
    EOF
  }
}

resource "null_resource" "azure_storage1_storage" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      az storage container delete --name "container_name"
    EOF
  }
}
```
