
# Terraform Infra
# Node: gcp_lb1
# Time: 1774805652

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
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name                  = "gcp-lb1"
  target                = google_compute_target_pool.gcp_lb1.self_link
  port_range            = "80"
  load_balancing_scheme = "INTERNAL"
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1"
  network = "default"
  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
}

resource "google_compute_autoscaler" "gcp_lb1" {
  name   = "gcp-lb1"
  target = google_compute_instance_group_manager.gcp_lb1.id

  autoscaling_policy {
    max_replicas    = 1
    min_replicas    = 1
    cooldown_period = 60
  }
}

resource "google_compute_disk" "gcp_lb1" {
  name  = "gcp-lb1"
  zone  = "us-central1-a"
  type  = "pd-standard"
  size  = 30
}

resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp-lb1"
  machine_type = "f1-micro"
  zone         = "us-central1-a"
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  attached_disk {
    source = google_compute_disk.gcp_lb1.self_link
  }
}

resource "google_compute_instance_group" "gcp_lb1" {
  name = "gcp-lb1"
  instances = [
    google_compute_instance.gcp_lb1.self_link,
  ]
}

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1"
  version {
    instance {
      machine_type = "f1-micro"
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

resource "google_compute_target_pool" "gcp_lb1" {
  name = "gcp-lb1"
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name                  = "gcp-lb
