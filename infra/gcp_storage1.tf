
# Terraform Infra
# Node: gcp_storage1
# Time: 1774805018

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1"
  zone  = "us-central1-a"
  size  = 50 // Changed from 75 to 50 GB
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = "instance-1"
  mode     = "READ_WRITE"
}

resource "google_compute_instance" "gcp_storage1" {
  name  = "gcp-storage-1"
  zone  = "us-central1-a"
  machine_type = "n1-standard-1" // Changed from n1-standard-4 to n1-standard-1
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  network_interface {
    network = "default"
  }
}
```
