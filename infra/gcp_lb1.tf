
# Terraform Infra
# Node: gcp_lb1
# Time: 1774813971

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

resource "google_compute_target_pool" "gcp_lb1" {
  name = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.id]
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_backend_bucket" "gcp_lb1" {
  name        = "gcp-lb1"
  description = "Secure backend bucket"
  bucket_name = "gcp-lb1-bucket"
}

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name       = "gcp-lb1"
  target     = google_compute_target_pool.gcp_lb1.self_link
  port_range = "80"
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1"
  network = "default"

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
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.gcp_lb1.self_link
}

resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp-lb1"
  machine_type = "f1-micro"
  zone         = "us-central1-a"
  network_interface {
    network = google_compute_network.gcp_lb1.self_link
    subnetwork = google_compute_subnetwork.gcp_lb1.self_link
  }
}

resource "google_compute_disk" "gcp_lb1" {
  name  = "gcp-lb1-disk"
  zone  = "us-central1-a"
  image = "debian-cloud/debian-9"
  size  = 10
}

resource "google_compute_attached_disk" "gcp_lb1" {
  disk     = google_compute_disk.gcp_lb1.self_link
  instance = google_compute_instance.gcp_lb1.self_link
}

resource "google_compute_instance_group" "gcp_lb1" {
  name = "gcp-lb1"
  instances = [google_compute_instance.gcp_lb1.self_link]
}

resource "google_compute_health_check" "gcp_lb1" {
  name               = "gcp-lb1"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_target_pool" "gcp_lb1" {
  name = "gcp-lb1"
  health_checks = [google_compute
