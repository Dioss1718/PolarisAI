
# Terraform Infra
# Node: azure_vm1
# Time: 1774587680

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
    action   = "DOWNSIZE_MEDIUM"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm resize --resource-group <resource_group> --name azure_vm1 --size Standard_DS2_v2
    EOF
  }
}
```

```terraform
resource "null_resource" "azure_vm1_update_tags" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
    action   = "DOWNSIZE_MEDIUM"
  }

  provisioner "local-exec" {
    command = <<EOF
      az vm update --resource-group <resource_group> --name azure_vm1 --set tags.utilization=10.50 --set tags.cost=136.50
    EOF
  }
}
```
