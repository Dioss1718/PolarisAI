
# Terraform Infra
# Node: gcp_lb1
# Time: 1774802248

```terraform
resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp-lb1"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = google_compute_network.vpc.self_link
  }
}

resource "google_compute_network" "vpc" {
  name                    = "vpc-network"
  auto_create_subnetworks = false
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = google_compute_network.vpc.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1_internal" {
  name    = "gcp-lb1-internal-firewall"
  network = google_compute_network.vpc.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1_compliance" {
  name    = "gcp-lb1-compliance-firewall"
  network = google_compute_network.vpc.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["10.0.0.0/8"]
}
```
