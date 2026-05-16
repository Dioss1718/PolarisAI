
# Terraform Infra
# Node: gcp_storage1
# Time: 1774805521

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}

resource "null_resource" "gcp_storage1_downsize_medium" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}
```
