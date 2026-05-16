
# Terraform Infra
# Node: gcp_storage1
# Time: 1774593017

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  metadata = {
    startup-script = <<-EOF
      #!/bin/bash
      echo "Instance resized to n1-standard-1"
    EOF
  }
}

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 30
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
      image = "debian-cloud/debian-11"
    }
  }

  metadata = {
    startup-script = <<-EOF
      #!/bin/bash
      echo "Instance resized to n1-standard-1"
    EOF
  }
}

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "null_resource" "resize" {
  provisioner "local-exec" {
    command = "gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1"
  }
}
```

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  metadata = {
    startup-script = <<-EOF
      #!/bin/bash
      echo "Instance resized to n1-standard-1"
    EOF
  }
}

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "null_resource" "resize" {
  provisioner "local-exec" {
    command = "gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1"
  }
}

resource "null_resource" "wait" {
  provisioner "local-exec" {
    command = "sleep 60"
  }
}
```

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
