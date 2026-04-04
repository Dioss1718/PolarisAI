
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808664

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  security_policy = google_compute_security_policy.gcp_lb1.name
  log_config {
    enable = true
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

resource "google_compute_security_policy" "gcp_lb1" {
  name = "gcp-lb1"
  rule {
    action = "ALLOW"
    priority = 1000
    match {
      config {
        src_ip_ranges = ["0.0.0.0/0"]
      }
    }
  }
  rule {
    action = "ALLOW"
    priority = 1001
    match {
      config {
        src_ip_ranges = ["35.235.240.0/20"]
      }
    }
  }
  rule {
    action = "ALLOW"
    priority = 1002
    match {
      config {
        src_ip_ranges = ["35.191.0.0/16"]
      }
    }
  }
  rule {
    action = "ALLOW"
    priority = 1003
    match {
      config {
        src_ip_ranges = ["130.211.0.0/22"]
      }
    }
  }
  rule {
    action = "ALLOW"
    priority = 1004
    match {
      config {
        src_ip_ranges = ["35.235.240.0/20"]
      }
    }
  }
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1"
  network = google_compute_network.gcp_lb1.name

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
  network       = google_compute_network.gcp_lb1.name
}

resource "google_compute_address" "gcp_lb1" {
  name         = "gcp-lb1"
  subnetwork   = google_compute_subnetwork.gcp_lb1.name
  address_type = "EXTERNAL"
}

resource "google_compute_target_http_proxy" "gcp_lb1" {
  name    = "gcp-lb1"
  url_map = google_compute_url_map.gcp_lb1.self_link
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_backend_service" "gcp_lb1_cost" {
  name          = "gcp-lb1-cost"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  security_policy = google_compute_security_policy.gcp_lb1.name
  log_config {
    enable = true
  }
}

resource
