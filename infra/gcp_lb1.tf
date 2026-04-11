
# Terraform Infra
# Node: gcp_lb1
# Time: 1775923814

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  security_policy = google_compute_security_policy.gcp_lb1.name
}

resource "google_compute_health_check" "gcp_lb1" {
  name               = "gcp-lb1"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1"
  version {
    instance {
      machine_type = "f1-micro"
    }
  }
}

resource "google_compute_security_policy" "gcp_lb1" {
  name = "gcp-lb1"
  rule {
    action = "ALLOW"
    priority = 1000
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["10.0.0.0/8"]
      }
    }
  }
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1"
  network = google_compute_network.gcp_lb1.name
  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
}

resource "google_compute_network" "gcp_lb1" {
  name                    = "gcp-lb1"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "gcp_lb1" {
  name          = "gcp-lb1"
  ip_cidr_range = "10.0.0.0/16"
  network       = google_compute_network.gcp_lb1.id
}
```
