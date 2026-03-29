
# Terraform Infra
# Node: gcp_lb1
# Time: 1774801631

```terraform
resource "google_compute_instance" "gcp_lb1" {
  // existing instance resource
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = google_compute_network.gcp_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp-lb1"]
}

resource "google_compute_network" "gcp_network" {
  name                    = "gcp-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "gcp_subnetwork" {
  name          = "gcp-subnetwork"
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.gcp_network.self_link
}

resource "google_compute_firewall_rule" "gcp_lb1_internal" {
  name    = "gcp-lb1-internal-firewall-rule"
  network = google_compute_network.gcp_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  source_ranges = ["10.0.1.0/24"]
}

resource "google_compute_firewall_rule" "gcp_lb1_public" {
  name    = "gcp-lb1-public-firewall-rule"
  network = google_compute_network.gcp_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp-lb1"]
}

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1-instance-group-manager"

  base_instance_name = "gcp-lb1"
  instance_template {
    // existing instance template resource
  }
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
  health_checks = [google_compute_health_check.gcp_lb1.self_link]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1-url-map"
  default_service = google_compute_backend_service.gcp_lb1.self_link

  host_rule {
    hosts        = ["lb1.example.com"]
    service      = google_compute_backend_service.gcp_lb1.self_link
  }

  path_matcher {
    name            = "lb1-path-matcher"
    default_service = google_compute_backend_service.gcp_lb1.self_link

    path_rule {
      service = google_compute_backend_service.gcp_lb1.self_link
    }
  }
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

resource "google_compute_global_forwarding_rule" "gcp_lb1_https" {
  name       = "gcp-lb1-global-forwarding-rule-https"
  target     = google_compute_target_http_proxy.gcp
