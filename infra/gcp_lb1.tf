
# Terraform Infra
# Node: gcp_lb1
# Time: 1775927334

```terraform
resource "google_compute_instance" "gcp_lb1" {
  // existing instance configuration
}

resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp-lb1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_internal" {
  name    = "gcp-lb1-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["gcp-lb1"]
  source_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal" {
  name    = "gcp-lb1-allow-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal_from_lb" {
  name    = "gcp-lb1-allow-internal-from-lb-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal_from_lb_internal" {
  name    = "gcp-lb1-allow-internal-from-lb-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_tags = ["gcp-lb1"]
  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal_from_lb_internal_from_lb" {
  name    = "gcp-lb1-allow-internal-from-lb-internal-from-lb-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_tags = ["gcp-lb1"]
  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal_from_lb_internal_from_lb_internal" {
  name    = "gcp-lb1-allow-internal-from-lb-internal-from-lb-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_tags = ["gcp-lb1"]
  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal_from_lb_internal_from_lb_internal_from_lb" {
  name    = "gcp-lb1-allow-internal-from-lb-internal-from-lb-internal-from-lb-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_tags = ["gcp-lb1"]
  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "gcp_lb1_allow_internal_from_lb_internal_from_lb_internal_from_lb_internal" {
  name    = "gcp-lb1-allow-internal-from-lb-internal-from-lb-internal-from-lb-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_tags = ["gcp-lb1"]
  target_tags = ["gcp-lb1"]
}

resource "google_compute_firewall" "
