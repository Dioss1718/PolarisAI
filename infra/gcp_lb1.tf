
# Terraform Infra
# Node: gcp_lb1
# Time: 1774809032

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute backend-services update gcp_lb1 \
      --region us-central1 \
      --access-policies ${google_compute_backend_service.gcp_lb1.access_policies[0].self_link} \
      --load-balancing-scheme INTERNAL_MANAGED \
      --update-quota 60.00 \
      --update-quota-unit "HOUR"
    EOF
  }

  connection {
    type        = "local"
    agent       = false
  }
}

resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp_lb1"
  region        = "us-central1"
  health_checks = [google_compute_health_check.gcp_lb1.self_link]

  access_policy {
    external_access = "ALLOW"
  }

  load_balancing_scheme = "INTERNAL_MANAGED"
  quota_unit           = "HOUR"
  quota               = 60.00
}

resource "google_compute_health_check" "gcp_lb1" {
  name               = "gcp_lb1"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
```
