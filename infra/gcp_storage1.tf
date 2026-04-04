
# Terraform Infra
# Node: gcp_storage1
# Time: 1774587695

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-4"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_disk" "gcp_storage1_disk" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1_disk_attachment" {
  disk     = google_compute_disk.gcp_storage1_disk.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "google_compute_firewall" "gcp_storage1_firewall" {
  name    = "gcp-storage-1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }
}

resource "google_compute_health_check" "gcp_storage1_healthcheck" {
  name               = "gcp-storage-1-healthcheck"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "22"
  }
}

resource "google_compute_backend_service" "gcp_storage1_backend" {
  name          = "gcp-storage-1-backend"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_storage1_healthcheck.self_link]
}

resource "google_compute_url_map" "gcp_storage1_urlmap" {
  name            = "gcp-storage-1-urlmap"
  default_service = google_compute_backend_service.gcp_storage1_backend.self_link
}

resource "google_compute_target_http_proxy" "gcp_storage1_proxy" {
  name    = "gcp-storage-1-proxy"
  url_map = google_compute_url_map.gcp_storage1_urlmap.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_storage1_forwarding_rule" {
  name       = "gcp-storage-1-forwarding-rule"
  target     = google_compute_target_http_proxy.gcp_storage1_proxy.self_link
  port_range = "80"
}

resource "google_compute_region_backend_service" "gcp_storage1_region_backend" {
  name          = "gcp-storage-1-region-backend"
  region        = "us-central1"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_storage1_healthcheck.self_link]
}

resource "google_compute_region_url_map" "gcp_storage1_region_urlmap" {
  name            = "gcp-storage-1-region-urlmap"
  default_service = google_compute_region_backend_service.gcp_storage1_region_backend.self_link
}

resource "google_compute_region_target_http_proxy" "gcp_storage1_region_proxy" {
  name    = "gcp-storage-1-region-proxy"
  url_map = google_compute_region_url_map.gcp_storage1_region_urlmap.self_link
}

resource "google_compute_region_global_forwarding_rule" "gcp_storage1_region_forwarding_rule" {
  name       = "gcp-storage-1-region-forwarding-rule"
  target     = google_compute_region_target_http_proxy.gcp_storage1_region_proxy.self_link
  port_range = "80"
}

resource "google_storage_bucket" "gcp_storage1_bucket" {
  name          = "gcp-storage-1-bucket"
  location      = "US"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket_iam
