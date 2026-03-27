
# Terraform Infra
# Node: gcp_lb1
# Time: 1774590149

```terraform
resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp-lb1"
  machine_type = "f1-micro"

  network_interface {
    network = "default"
  }

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  metadata = {
    startup-script = <<-EOF
      #! /bin/bash
      apt-get update && apt-get install -y apache2
      echo "Hello from GCP!" > /var/www/html/index.html
      service apache2 start
    EOF
  }
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_health_check" "gcp_lb1" {
  name               = "gcp-lb1-healthcheck"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1-backend"
  port_name     = "http"
  protocol      = "HTTP"
  health_checks = [google_compute_health_check.gcp_lb1.self_link]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }
}

resource "google_compute_instance_group_manager" "gcp_lb1" {
  name = "gcp-lb1-igm"

  base_instance_name = "gcp-lb1"
  instance_template  = google_compute_instance_template.gcp_lb1.self_link
  target_size        = 1
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1-template"
  machine_type  = "f1-micro"
  can_ip_forward = false

  disk {
    source_image = "debian-cloud/debian-9"
  }

  network_interface {
    network = "default"
  }

  metadata = {
    startup-script = <<-EOF
      #! /bin/bash
      apt-get update && apt-get install -y apache2
      echo "Hello from GCP!" > /var/www/html/index.html
      service apache2 start
    EOF
  }
}
```
