
# Terraform Infra
# Node: gcp_storage1
# Time: 1774805695

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 14 // GB
  type  = "pd-standard"
  labels = {
    environment = "production"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = "instance-1"
  mode     = "READ_WRITE"
}

resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
      size  = 14 // GB
    }
  }
  network_interface {
    network = "default"
  }
  metadata = {
    startup-script = <<-EOF
      #! /bin/bash
      apt-get update -y
      apt-get install -y nginx
      service nginx start
    EOF
  }
}
```
