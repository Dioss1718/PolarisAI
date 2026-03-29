
# Terraform Infra
# Node: azure_storage1
# Time: 1774815639

```terraform
resource "null_resource" "azure_storage1" {
  triggers = {
    node_id = "azure_storage1"
    action  = "TERMINATE_FORCE"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      terraform taint -force ${node_id}
      terraform apply -auto-approve
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
      terraform state rm ${node_id}
      terraform apply -auto-approve
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
      terraform state mv ${node_id} deleted.${node_id}
      terraform apply -auto-approve
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
      terraform state rm ${node_id}
      terraform state rm ${node_id}_exposure
      terraform apply -auto-approve
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
      terraform state rm ${node_id}
      terraform state rm ${node_id}_exposure
      terraform state rm ${node_id}_utilization
      terraform state rm ${node_id}_cost
      terraform apply -auto-approve
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
      terraform state rm ${node_id}
      terraform state rm ${node_id}_exposure
      terraform state rm ${node_id}_utilization
      terraform state rm ${node_id}_cost
      terraform state rm ${node_id}_exposure_public
      terraform state rm ${node_id}_exposure_private
      terraform apply -auto-approve
    EOF
  }
}
```
