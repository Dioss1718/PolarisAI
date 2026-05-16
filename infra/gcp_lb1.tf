
# Terraform Infra
# Node: gcp_lb1
# Time: 1774801421

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name     = "gcp-lb1"
  protocol = "HTTP"

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  health_checks = [google_compute_health_check.gcp_lb1.name]
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

  base_instance_name = "gcp-lb1"
  instance_template  = google_compute_instance_template.gcp_lb1.self_link
  target_size        = 1
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1"
  machine_type  = "f1-micro"
  disk {
    source_image = "debian-cloud/debian-9"
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

  source_ranges = ["10.0.0.0/8"]
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

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name       = "gcp-lb1"
  target     = google_compute_target_http_proxy.gcp_lb1.self_link
  port_range = "80"
}

resource "google_compute_region_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  region        = "us-central1"
  load_balancing_scheme = "INTERNAL"
  health_checks = [google_compute_health_check.gcp_lb1.name]
}

resource "google_compute_region_network_endpoint_group" "gcp_lb1" {
  name          = "gcp-lb1"
  region        = "us-central1"
  network       = google_compute_network.gcp_lb1.self_link
  subnet        = google_compute_subnetwork.gcp_lb1.self_link
  default_port  = 80
}

resource "google_compute_subnetwork" "gcp_lb1" {
  name          = "gcp-lb1"
  region        = "us-central1"
  network       = google_compute_network.gcp_lb1.self_link
  ip_cidr_range = "10.0.0.0/24"
}

resource "google_compute_region_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  region        = "us-central1"
  load_balancing_scheme = "INTERNAL"
  health_checks = [google_compute_health_check.gcp_lb1.name]
