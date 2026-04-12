
# Terraform Infra
# Node: gcp_storage1
# Time: 1775965773

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-4
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-2
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-1
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-8
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-4
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-2
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-1
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-8
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp_storage1 --zone=your-zone --machine-type=n1-standard-4
    EOF
  }
}
```

```terraform
resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      g
