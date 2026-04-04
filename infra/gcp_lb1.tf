
# Terraform Infra
# Node: gcp_lb1
# Time: 1774596431

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

resource "google_compute_backend_bucket" "gcp_lb1" {
  name        = "gcp-lb1"
  bucket_name = "gcp-lb1-bucket"
}

resource "google_compute_global_forwarding_rule" "gcp_lb1" {
  name       = "gcp-lb1"
  target     = google_compute_target_pool.gcp_lb1.self_link
  port_range = "80"
}

resource "google_compute_region_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  load_balancing_scheme = "INTERNAL_MANAGED"
  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_region_health_check" "gcp_lb1" {
  name               = "gcp-lb1"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_url_map" "gcp_lb1" {
  name            = "gcp-lb1"
  default_service = google_compute_region_backend_service.gcp_lb1.self_link
}

resource "google_compute_region_backend_service_iam_member" "gcp_lb1" {
  role    = "roles/compute.regionBackendServiceUser"
  member  = "serviceAccount:service-1234567890@compute-system.iam.gserviceaccount.com"
  service = google_compute_region_backend_service.gcp_lb1.self_link
}

resource "google_compute_region_backend_service_iam_policy" "gcp_lb1" {
  policy_data = data.google_iam_policy_admin_view.policy_data
  service     = google_compute_region_backend_service.gcp_lb1.self_link
}

data "google_iam_policy_admin_view" "policy_data" {
  depends_on = [google_compute_region_backend_service.gcp_lb1]
}

resource "google_compute_region_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  load_balancing_scheme = "INTERNAL_MANAGED"
  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
  cost = 63.00
  security_policy = google_compute_security_policy.gcp_lb1.self_link
}

resource "google_compute_security_policy" "gcp_lb1" {
  name = "gcp-lb1"
  rule {
    action = "ALLOW"
    priority = 1000
    match {
      versioned_expr = "SRC_IP"
      config {
        src_service = "ALL_SERVICES"
      }
    }
  }
  rule {
    action = "
