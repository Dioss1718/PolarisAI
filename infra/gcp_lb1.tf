
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808511

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
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
    instance {
      zone = "us-central1-a"
      machine_type = "f1-micro"
      disk {
        image = "debian-cloud/debian-9"
      }
    }
  }
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["gcp-lb1"]
}

resource "google_compute_network" "gcp_lb1" {
  name                    = "gcp-lb1"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "gcp_lb1" {
  name          = "gcp-lb1"
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.gcp_lb1.self_link
  region        = "us-central1"
}

resource "google_compute_address" "gcp_lb1" {
  name         = "gcp-lb1"
  address_type = "EXTERNAL"
  region       = "us-central1"
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_target_http_proxy" "gcp_lb1" {
  name    = "gcp-lb1"
  url_map = google_compute_url_map.gcp_lb1.self_link
}

resource "google_compute_target_https_proxy" "gcp_lb1" {
  name             = "gcp-lb1"
  url_map          = google_compute_url_map.gcp_lb1.self_link
  ssl_certificates = [google_compute_ssl_certificate.gcp_lb1.self_link]
}

resource "google_compute_ssl_certificate" "gcp_lb1" {
  name             = "gcp-lb1"
  private_key      = file("~/.ssh/gcp-lb1-key.pem")
  certificate      = file("~/.ssh/gcp-lb1-cert.pem")
}

resource "google_compute_backend_service" "gcp_lb1_https" {
  name          = "gcp-lb1-https"
  port_name     = "https"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_url_map" "gcp_lb1_https" {
  name            = "gcp-lb1-https"
  default_service = google_compute_backend_service.gcp_lb1_https.self_link
}

resource "google_compute_target_http_proxy" "gcp_lb1_https" {
  name    = "gcp-lb1-https"
  url_map = google_compute_url_map.gcp_lb1_https.self_link
}

resource "google_compute_target_https_proxy" "gcp_lb1_https
