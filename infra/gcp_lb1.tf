
# Terraform Infra
# Node: gcp_lb1
# Time: 1774588563

```terraform
resource "google_compute_backend_service" "gcp_lb1" {
  name          = "gcp-lb1"
  health_checks = [google_compute_health_check.gcp_lb1.name]

  backend {
    group = google_compute_instance_group_manager.gcp_lb1.instance_group
  }

  load_balancing_scheme = "INTERNAL"
  connection_draining   = true
  enable_cdn            = false
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

  base_instance_name = "gcp-lb1"
  instance_template  = google_compute_instance_template.gcp_lb1.self_link
  target_size        = 1
}

resource "google_compute_instance_template" "gcp_lb1" {
  name          = "gcp-lb1"
  machine_type  = "f1-micro"
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

resource "google_compute_firewall" "gcp_lb1-allow-internal" {
  name    = "gcp-lb1-allow-internal"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1-allow-internal-icmp" {
  name    = "gcp-lb1-allow-internal-icmp"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "icmp"
  }
  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1-allow-internal-udp" {
  name    = "gcp-lb1-allow-internal-udp"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "udp"
    ports    = ["80"]
  }
  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1-allow-internal-tcp" {
  name    = "gcp-lb1-allow-internal-tcp"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1-allow-internal-icmpv6" {
  name    = "gcp-lb1-allow-internal-icmpv6"
  network = google_compute_network.gcp_lb1.self_link

  allow {
    protocol = "icmp"
  }
  source_ranges = ["2001:0db8::/32"]
}

resource "google_compute_firewall" "gcp_lb1-allow
