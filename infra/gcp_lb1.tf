
# Terraform Infra
# Node: gcp_lb1
# Time: 1774599274

```terraform
resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = google_compute_network.gcp_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_network" "gcp_network" {
  name                    = "gcp-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "gcp_subnetwork" {
  name          = "gcp-subnetwork"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.gcp_network.self_link
}

resource "google_compute_instance" "gcp_instance" {
  name         = "gcp-instance"
  machine_type = "f1-micro"
  zone         = "us-central1-a"
  network_interface {
    network = google_compute_network.gcp_network.self_link
    subnetwork = google_compute_subnetwork.gcp_subnetwork.self_link
  }
}

resource "google_compute_health_check" "gcp_healthcheck" {
  name               = "gcp-healthcheck"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_backend_service" "gcp_backend" {
  name          = "gcp-backend"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_healthcheck.self_link]
}

resource "google_compute_url_map" "gcp_urlmap" {
  name            = "gcp-urlmap"
  default_service = google_compute_backend_service.gcp_backend.self_link
}

resource "google_compute_target_http_proxy" "gcp_proxy" {
  name    = "gcp-proxy"
  url_map = google_compute_url_map.gcp_urlmap.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_forwarding" {
  name       = "gcp-forwarding"
  target     = google_compute_target_http_proxy.gcp_proxy.self_link
  port_range = "80"
}

resource "google_compute_region_backend_service" "gcp_region_backend" {
  name          = "gcp-region-backend"
  region        = "us-central1"
  load_balancing_scheme = "INTERNAL"
  health_checks = [google_compute_health_check.gcp_healthcheck.self_link]
}

resource "google_compute_region_url_map" "gcp_region_urlmap" {
  name            = "gcp-region-urlmap"
  default_service = google_compute_region_backend_service.gcp_region_backend.self_link
}

resource "google_compute_region_target_http_proxy" "gcp_region_proxy" {
  name    = "gcp-region-proxy"
  url_map = google_compute_region_url_map.gcp_region_urlmap.self_link
}

resource "google_compute_region_global_forwarding_rule" "gcp_region_forwarding" {
  name       = "gcp-region-forwarding"
  target     = google_compute_region_target_http_proxy.gcp_region_proxy.self_link
  port_range = "80"
}

resource "google_compute_region_backend_service" "gcp_region_backend" {
  name          = "gcp-region-backend"
  region        = "us-central1"
  load_balancing_scheme = "INTERNAL"
  health_checks = [google_compute_health_check.gcp_healthcheck.self_link]
}

resource "google_compute_region_url_map" "gcp_region_urlmap" {
  name            = "gcp-region-urlmap"
  default_service = google_compute_region_backend_service.gcp_region_backend.self_link
}

resource "google_compute_region_target_http_proxy" "
