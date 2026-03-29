
# Terraform Infra
# Node: gcp_lb1
# Time: 1774802152

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1-backend-service"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.id]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_health_check" "gcp_lb1" {
  name               = "gcp-lb1-health-check"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1-instance-group-manager"

  base_instance_name = "gcp-lb1-instance"
  instance_template  = google_compute_instance_template.gcp_lb1.self_link
  target_size        = 1
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1-instance-template"
  machine_type  = "f1-micro"
  disk {
    source_image = "debian-cloud/debian-9"
  }
  network_interface {
    network = google_compute_network.gcp_lb1.self_link
  }
}

resource "google_compute_network" "gcp_lb1" {
  name                    = "gcp-lb1-network"
  auto_create_subnetworks = false
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
}
```
