
# Terraform Infra
# Node: gcp_storage1
# Time: 1774546411

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

  network_interface {
    network = google_compute_network.default.self_link
  }
}

resource "google_compute_network" "default" {
  name                    = "default-network"
  auto_create_subnetworks = false
}

resource "google_compute_firewall" "default" {
  name    = "default-firewall"
  network = google_compute_network.default.self_link

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }
}

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 10
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}
```
