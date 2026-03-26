
# Terraform Infra
# Node: gcp_storage1
# Time: 1774543492

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_disk" "gcp_storage1_disk" {
  name  = "gcp-storage1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
  labels = {
    environment = "dev"
  }
}

resource "google_compute_attached_disk" "gcp_storage1_disk_attachment" {
  disk     = google_compute_disk.gcp_storage1_disk.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "google_compute_firewall" "gcp_storage1_firewall" {
  name    = "gcp-storage1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }
}

resource "google_compute_health_check" "gcp_storage1_healthcheck" {
  name               = "gcp-storage1-healthcheck"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "22"
  }
}

resource "google_compute_backend_service" "gcp_storage1_backend" {
  name          = "gcp-storage1-backend"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_storage1_healthcheck.self_link]
}

resource "google_compute_url_map" "gcp_storage1_urlmap" {
  name            = "gcp-storage1-urlmap"
  default_service = google_compute_backend_service.gcp_storage1_backend.self_link
}

resource "google_compute_target_http_proxy" "gcp_storage1_proxy" {
  name    = "gcp-storage1-proxy"
  url_map = google_compute_url_map.gcp_storage1_urlmap.self_link
}

resource "google_compute_global_forwarding_rule" "gcp_storage1_forwarding_rule" {
  name       = "gcp-storage1-forwarding-rule"
  target     = google_compute_target_http_proxy.gcp_storage1_proxy.self_link
  port_range = "80"
}

resource "google_compute_instance_group_manager" "gcp_storage1_igm" {
  name        = "gcp-storage1-igm"
  base_instance_name = "gcp-storage1"
  zone         = "us-central1-a"

  version {
    instance {
      zone = "us-central1-a"
      machine_type = "n1-standard-1"
      boot_disk {
        initialize_params {
          image = "debian-cloud/debian-10"
        }
      }
      network_interface {
        network = "default"
      }
    }
  }
}

resource "google_compute_autoscaler" "gcp_storage1_autoscaler" {
  name   = "gcp-storage1-autoscaler"
  zone   = "us-central1-a"
  target = google_compute_instance_group_manager.gcp_storage1_igm.self_link

  autoscaling_policy {
    max_replicas    = 1
    min_replicas    = 1
    cooldown_period = 60
  }
}

resource "google_compute_region_autoscaling_policy" "gcp_storage1_region_autoscaling_policy" {
  name  = "gcp-storage1-region-autoscaling-policy"
  region = "us-central1"

  autoscaling_policy {
    max_replicas    = 1
    min_replicas    = 1
    cooldown_period = 60
  }
}

resource "
