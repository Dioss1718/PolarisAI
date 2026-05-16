
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808566

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  security_policy = google_compute_security_policy.gcp_lb1.name
  timeout_sec     = 10

  log_config {
    enable = true
  }
}

resource "google_compute_security_policy" "gcp_lb1" {
  name        = "gcp-lb1"
  display_name = "gcp-lb1"

  rule {
    action   = "ALLOW"
    priority = 1000
    match {
      versioned_expr = "SRC_IP"
      config {
        src_service = "all_services"
      }
    }
  }
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
    instance_template = google_compute_instance_template.gcp_lb1.self_link
    name              = "primary"
  }
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1"
  machine_type  = "f1-micro"
  can_ip_forward = false

  disk {
    source_image = "debian-cloud/debian-9"
    auto_delete  = true
  }

  network_interface {
    network = google_compute_network.gcp_lb1.self_link
  }
}

resource "google_compute_network" "gcp_lb1" {
  name                    = "gcp-lb1"
  auto_create_subnetworks = false
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
}
```
