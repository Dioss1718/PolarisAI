
# Terraform Infra
# Node: gcp_lb1
# Time: 1774601254

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
    network = "default"
  }
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_internal" {
  name    = "gcp-lb1-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["10.0.0.0/8"]
  target_tags   = ["gcp-lb1"]
}

resource "google_compute_health_check" "gcp_lb1" {
  name               = "gcp-lb1-health-check"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1-backend-service"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1-url-map"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_target_http_proxy" "gcp_lb1" {
  name    = "gcp-lb1-target-http-proxy"
  url_map = google_compute_url_map.gcp_lb1.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name       = "gcp-lb1-global-forwarding-rule"
  target     = google_compute_target_http_proxy.gcp_lb1.self_link
  port_range = "80"
}

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1-instance-group-manager"

  base_instance_name = "gcp-lb1-instance"
  instance_template  = google_compute_instance_template.gcp_lb1.self_link
  target_size        = 1
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1-instance-template"
  machine_type  = "f1-micro"
  disk {
    source_image = "debian-cloud/debian-9"
  }

  network_interface {
    network = "default"
  }
}
```
