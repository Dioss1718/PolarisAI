
# Terraform Infra
# Node: gcp_storage1
# Time: 1775966276

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
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone us-central1-a --machine-type n1-standard-4"
  }
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone us-central1-a --machine-type n1-standard-1"
  }
}
```
