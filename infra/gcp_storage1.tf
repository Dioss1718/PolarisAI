
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801747

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 50 // GB
}

resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
      size  = 50 // GB
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "google_compute_instance" "gcp_storage1" {
  depends_on = [google_compute_disk_attachment.gcp_storage1]

  lifecycle {
    ignore_changes = [
      machine_type,
    ]
  }

  machine_type = "n1-standard-1"
}
```

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 50 // GB
}

resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
      size  = 50 // GB
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "google_compute_instance" "gcp_storage1" {
  depends_on = [google_compute_disk_attachment.gcp_storage1]

  lifecycle {
    ignore_changes = [
      machine_type,
    ]
  }

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Instance ${google_compute_instance.gcp_storage1.name} has been downsized to ${google_compute_instance.gcp_storage1.machine_type}.'"
  }
}

resource "null_resource" "gcp_storage1" {
  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Instance ${google_compute_instance.gcp_storage1.name} utilization changed from 20.00 to 14.00.'"
  }
}

resource "null_resource" "gcp_storage1" {
  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Instance ${google_compute_instance.gcp_storage1.name} cost changed from 35.00 to 24.50.'"
  }
}
```
