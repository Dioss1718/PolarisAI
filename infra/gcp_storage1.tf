
# Terraform Infra
# Node: gcp_storage1
# Time: 1774809042

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1"
  }
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-4"
  }
}

resource "null_resource" "gcp_storage1_downsize_medium" {
  triggers = {
    instance_size = "n1-standard-4"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-2"
  }
}

resource "null_resource" "gcp_storage1_downsize_medium_final" {
  triggers = {
    instance_size = "n1-standard-2"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1"
  }
}
```
