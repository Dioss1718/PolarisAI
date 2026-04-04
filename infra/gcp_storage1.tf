
# Terraform Infra
# Node: gcp_storage1
# Time: 1774802157

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_disk" "gcp_storage1_disk" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 10
  type  = "pd-standard"
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1
      gcloud compute disks update gcp-storage-1-disk --zone us-central1-a --size 10
    EOF
  }
}
```
