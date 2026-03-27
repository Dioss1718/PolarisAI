
# Terraform Infra
# Node: gcp_storage1
# Time: 1774592051

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
}

resource "google_compute_disk" "gcp_storage1_disk" {
  name  = "gcp-storage1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1_disk_attach" {
  disk     = google_compute_disk.gcp_storage1_disk.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "google_compute_firewall" "gcp_storage1_firewall" {
  name    = "gcp-storage1-firewall"
  network = google_compute_network.gcp_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }
}

resource "google_compute_network" "gcp_network" {
  name                    = "gcp-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "gcp_subnetwork" {
  name          = "gcp-subnetwork"
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.gcp_network.self_link
  region        = "us-central1"
}

resource "google_compute_instance_group" "gcp_instance_group" {
  name = "gcp-instance-group"
}

resource "google_compute_instance_template" "gcp_instance_template" {
  name          = "gcp-instance-template"
  machine_type  = "n1-standard-1"
  disk {
    source_image = "debian-cloud/debian-9"
  }
  network_interface {
    network = google_compute_network.gcp_network.self_link
  }
}

resource "google_compute_health_check" "gcp_health_check" {
  name               = "gcp-health-check"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "22"
  }
}

resource "google_compute_target_pool" "gcp_target_pool" {
  name = "gcp-target-pool"
}

resource "google_compute_backend_service" "gcp_backend_service" {
  name          = "gcp-backend-service"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_health_check.self_link]
}

resource "google_compute_url_map" "gcp_url_map" {
  name            = "gcp-url-map"
  default_service = google_compute_backend_service.gcp_backend_service.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_global_forwarding_rule" {
  name                  = "gcp-global-forwarding-rule"
  target                = google_compute_target_pool.gcp_target_pool.self_link
  port_range            = "80"
  load_balancing_scheme = "EXTERNAL"
}

resource "google_compute_target_http_proxy" "gcp_target_http_proxy" {
  name    = "gcp-target-http-proxy"
  url_map = google_compute_url_map.gcp_url_map.self_link
}

resource "google_compute_ssl_policy" "gcp_ssl_policy" {
  name            = "gcp-ssl-policy"
  min_tls_version = "TLS_1_2"
}

resource "google_compute_backend_service_ssl_policy" "gcp_backend_service_ssl_policy" {
  backend_service = google_compute_backend_service.gcp_backend_service.self_link
  ssl_policy      = google_compute_ssl_policy.gcp_ssl_policy.self_link
}

resource "google_compute_target_ssl_proxy" "gcp_target_ssl_proxy" {
  name             = "gcp-target-ssl-proxy"
  url_map          = google_compute_url_map.gcp_url_map.self_link
