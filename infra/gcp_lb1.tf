
# Terraform Infra
# Node: gcp_lb1
# Time: 1774588000

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.self_link]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  security_policy = google_compute_security_policy.gcp_lb1.self_link
}

resource "google_compute_security_policy" "gcp_lb1" {
  name        = "gcp-lb1"
  display_name = "gcp-lb1"

  rule {
    action   = "ALLOW"
    priority = 1000
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["0.0.0.0/0"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1001
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["35.235.240.0/20"]
      }
    }
  }

  rule {
    action   = "DENY"
    priority = 1002
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["0.0.0.0/0"]
      }
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

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1"

  base_instance_name = "gcp-lb1"
  instance_template  = google_compute_instance_template.gcp_lb1.self_link
  target_size        = 1
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1"
  machine_type  = "f1-micro"
  can_ip_forward = false

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
}

resource "google_compute_target_http_proxy" "gcp_lb1" {
  name    = "gcp-lb1"
  url_map = google_compute_url_map.gcp_lb1.self_link
}

resource "google_compute_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_backend_service.gcp_lb1.self_link
}

resource "google_compute_backend_bucket" "gcp_lb1" {
  name          = "gcp-lb1"
  bucket_name   = "gcp-lb1-bucket"
  enable_cdn    = true
  cache_mode    = "USE_ORIGIN_HEADERS"
  serve_while_stale = 3600
}

resource "google_compute_backend_config" "gcp_lb1" {
  name    = "gcp-lb1"
  cdn_cache_key = "gcp-lb1"
}

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name       = "gcp-lb1"
  target     = google_compute_target_http_proxy.gcp_lb1.self_link
