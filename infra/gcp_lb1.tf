
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808832

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  port_name     = "http"
  protocol      = "HTTP"
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
    action   = "ALLOW"
    priority = 1002
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["35.191.0.0/16"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1003
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["130.211.0.0/22"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1004
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["209.85.172.0/19"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1005
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["35.235.240.0/20"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1006
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["35.191.0.0/16"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1007
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["130.211.0.0/22"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1008
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["209.85.172.0/19"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1009
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["35.235.240.0/20"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1010
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["35.191.0.0/16"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1011
    match {
      versioned_expr = "SRC_IP"
      config {
        src_ip_ranges = ["130.211.0.0/22"]
      }
    }
  }

  rule {
    action   = "ALLOW"
    priority = 1012
    match {
