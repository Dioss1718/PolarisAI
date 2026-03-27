
# Terraform Infra
# Node: gcp_storage1
# Time: 1774587587

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1
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
      gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}

resource "google_compute_instance" "gcp_storage1_downsize" {
  name         = "gcp-storage-1-downsize"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}
```
