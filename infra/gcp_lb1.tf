
# Terraform Infra
# Node: gcp_lb1
# Time: 1775929195

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.id]

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
      machine_type = "f1-micro"
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
  private_key      = file("~/.ssh/gcp-lb1-key")
  certificate      = file("~/.ssh/gcp-lb1-cert")
  certificate_chain = file("~/.ssh/gcp-lb1-chain")
}

resource "google_compute_backend_bucket" "gcp_lb1" {
  name        = "gcp-lb1"
  bucket_name = "gcp-lb1-bucket"
}

resource "google_compute_region_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  region        = "us-central1"
  load_balancing_scheme = "INTERNAL_MANAGED"
  backend {
    group = google_compute_region_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_region_health_check" "gcp_lb1" {
  name               = "gcp-lb1"
  region             = "us-central1"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1"
  region = "us-central1
