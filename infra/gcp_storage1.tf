
# Terraform Infra
# Node: gcp_storage1
# Time: 1774595979

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  metadata = {
    startup-script = <<-EOF
      #!/bin/bash
      echo "Instance resized to n1-standard-1"
    EOF
  }
}

resource "google_compute_disk" "gcp_storage1_disk" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1_disk_attach" {
  disk     = google_compute_disk.gcp_storage1_disk.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}
```
