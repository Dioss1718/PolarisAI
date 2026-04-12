
# Terraform Infra
# Node: gcp_lb1
# Time: 1774600894

```terraform
resource "google_compute_instance" "gcp_lb1" {
  // existing instance properties...

  network_interface {
    network = "your-network-name"
    access_config {
      // existing access config...
    }
  }

  // existing instance properties...
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = "your-network-name"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  source_ranges = ["10.0.0.0/8"] // internal IP range
}

resource "google_compute_firewall" "gcp_lb1-egress" {
  name    = "gcp-lb1-egress-firewall"
  network = "your-network-name"

  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }

  allow {
    protocol = "udp"
    ports    = ["0-65535"]
  }

  allow {
    protocol = "icmp"
  }
}
```
