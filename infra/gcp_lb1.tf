
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808453

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  load_balancing_scheme = "INTERNAL"
  connection_draining   = true
  timeout_sec           = 10
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
    auto_delete  = true
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

resource "google_compute_address" "gcp_lb1" {
  name         = "gcp-lb1"
  address_type = "INTERNAL"
  subnetwork   = google_compute_subnetwork.gcp_lb1.self_link
}

resource "google_compute_subnetwork" "gcp_lb1" {
  name          = "gcp-lb1"
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.gcp_lb1.self_link
}

resource "google_compute_backend_service" "gcp_lb1_cost" {
  name          = "gcp-lb1-cost"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  load_balancing_scheme = "INTERNAL"
  connection_draining   = true
  timeout_sec           = 10
}

resource "google_compute_resource_policy" "gcp_lb1" {
  name = "gcp-lb1"

  recurring_schedule {
    recurring_schedule_spec {
      interval = "1 0 * * *"
    }
  }

  policy {
    cost  = 63.00
    usage = "EXTERMINATE"
  }
}
```
