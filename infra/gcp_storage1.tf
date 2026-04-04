
# Terraform Infra
# Node: gcp_storage1
# Time: 1774805225

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = google_compute_network.default.self_link
  }
}

resource "google_compute_network" "default" {
  name                    = "default-network"
  auto_create_subnetworks = "true"
}

resource "google_compute_firewall" "default" {
  name    = "default-firewall"
  network = google_compute_network.default.self_link

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }
}

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 10
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}
```

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = google_compute_network.default.self_link
  }
}

resource "google_compute_network" "default" {
  name                    = "default-network"
  auto_create_subnetworks = "true"
}

resource "google_compute_firewall" "default" {
  name    = "default-firewall"
  network = google_compute_network.default.self_link

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }
}

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 10
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "null_resource" "remediation" {
  provisioner "local-exec" {
    command = "gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1"
  }
}
```
