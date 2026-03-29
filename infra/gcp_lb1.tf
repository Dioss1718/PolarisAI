
# Terraform Infra
# Node: gcp_lb1
# Time: 1774817726

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.id]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  log_config {
    enable = true
  }

  security_policy = google_compute_security_policy.gcp_lb1.id

  dynamic "timeout" {
    for_each = [60]
    content {
      seconds = timeout.value
    }
  }

  dynamic "iap" {
    for_each = [true]
    content {
      oauth2_client_id = "your_client_id"
      oauth2_client_secret = "your_client_secret"
      enabled = iap.value
    }
  }

  dynamic "session_affinity" {
    for_each = ["CLIENT_IP"]
    content {
      session_affinity = session_affinity.value
    }
  }

  dynamic "load_balancing_scheme" {
    for_each = ["INTERNAL"]
    content {
      load_balancing_scheme = load_balancing_scheme.value
    }
  }

  dynamic "circuit_breakers" {
    for_each = [true]
    content {
      http {
        requests = {
          max_requests = 100
          max_elapsed_time = 60
        }
      }
    }
  }

  dynamic "outlier_detection" {
    for_each = [true]
    content {
      http {
        base_ejection_time = "300s"
        max_ejection_percent = 10
        min_health_percent = 99
        request_volume_threshold = 10
      }
    }
  }

  dynamic "locality_lb_settings" {
    for_each = [true]
    content {
      enable_cdn = true
      fallback_value = "0.0.0.0"
    }
  }

  dynamic "custom_request_headers" {
    for_each = [true]
    content {
      append_http_header {
        header_name  = "X-Custom-Header"
        header_value = "Custom-Value"
      }
    }
  }

  dynamic "custom_response_headers" {
    for_each = [true]
    content {
      append_http_header {
        header_name  = "X-Custom-Header"
        header_value = "Custom-Value"
      }
    }
  }

  dynamic "security_settings" {
    for_each = [true]
    content {
      ssl_certificates = [google_compute_ssl_certificate.gcp_lb1.id]
      client_tls_policy = google_compute_client_tls_policy.gcp_lb1.id
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      allow_rules {
        protocol = "TCP"
        ports    = ["80"]
      }
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      allow_rules {
        protocol = "UDP"
        ports    = ["53"]
      }
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      allow_rules {
        protocol = "ICMP"
        ports    = ["8"]
      }
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      deny_rules {
        protocol = "TCP"
        ports    = ["22"]
      }
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      deny_rules {
        protocol = "UDP"
        ports    = ["123"]
      }
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      deny_rules {
        protocol = "ICMP"
        ports    = ["3"]
      }
    }
  }

  dynamic "access_control" {
    for_each = [true]
    content {
      deny_rules
