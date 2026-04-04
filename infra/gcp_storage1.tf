
# Terraform Infra
# Node: gcp_storage1
# Time: 1774588007

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1"
  zone  = "us-central1-a"
  size  = 50 // GB
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = "gcp-instance-1"
  mode     = "READ_WRITE"
}

resource "google_compute_instance" "gcp_instance1" {
  name         = "gcp-instance-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
      size  = 50 // GB
    }
  }

  // Downsize medium instance
  // Original size: 1 vCPU, 3.75 GB RAM
  // New size: 1 vCPU, 1.7 GB RAM
  metadata = {
    // Downsize medium instance
    // Original size: 1 vCPU, 3.75 GB RAM
    // New size: 1 vCPU, 1.7 GB RAM
    "google-compute-enable-confidential-vm" = "true"
    "google-vm-confidential-mode"           = "true"
    "google-vm-disable-serial-console"      = "true"
  }
}
```
