
# Terraform Infra
# Node: azure_storage1
# Time: 1774546394

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      az storage container delete --name azure_storage1 --resource-group <resource_group_name>
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
      az storage container update --name azure_storage1 --public-access off --resource-group <resource_group_name>
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
      az storage container update --name azure_storage1 --metrics enabled --resource-group <resource_group_name>
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
      az storage container update --name azure_storage1 --usage-type Hot --resource-group <resource_group_name>
    EOF
  }
}
```
