
# Terraform Infra
# Node: azure_storage1
# Time: 1774809716

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm stop --ids $(az vm list --query "[?name=='${var.node_id}'].id" --output tsv)
      az vm deallocate --ids $(az vm list --query "[?name=='${var.node_id}'].id" --output tsv)
      az vm delete --ids $(az vm list --query "[?name=='${var.node_id}'].id" --output tsv) --yes
    EOF
  }
}
```

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az storage container delete --name ${var.node_id} --yes
    EOF
  }
}
```

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az storage account delete --name ${var.node_id} --yes
    EOF
  }
}
```
