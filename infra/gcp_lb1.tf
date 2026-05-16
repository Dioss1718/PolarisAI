
# Terraform Infra
# Node: gcp_lb1
# Time: 1774592368

```terraform
resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp_lb1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
}

resource "google_compute_firewall" "gcp_lb1_internal" {
  name    = "gcp_lb1-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_tags = ["gcp_lb1"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional" {
  name    = "gcp_lb1-internal-additional-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_2" {
  name    = "gcp_lb1-internal-additional-2-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["172.16.0.0/12"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_3" {
  name    = "gcp_lb1-internal-additional-3-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["192.168.0.0/16"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_4" {
  name    = "gcp_lb1-internal-additional-4-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["fd00::/8"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_5" {
  name    = "gcp_lb1-internal-additional-5-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["2001:db8::/32"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_6" {
  name    = "gcp_lb1-internal-additional-6-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["fc00::/7"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_7" {
  name    = "gcp_lb1-internal-additional-7-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  target_tags = ["gcp_lb1"]
  source_ranges = ["2001:10::/32"]
}

resource "google_compute_firewall" "gcp_lb1_internal_additional_8" {
  name    = "gcp_lb1-internal-additional-8-firewall"
  network = "default"

  allow {
    protocol = "tcp"
