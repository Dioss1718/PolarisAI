
# Terraform Infra
# Node: gcp_storage1
# Time: 1774817738

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 50 // GB
  type  = "pd-standard"
  labels = {
    env = "dev"
  }
}

resource "google_compute_disk" "gcp_storage1_downsized" {
  name  = "gcp-storage1-downsized"
  zone  = "us-central1-a"
  size  = 40 // GB (downsized by 10 GB)
  type  = "pd-standard"
  labels = {
    env = "dev"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.id
  instance = "instance-1"
  mode     = "READ_WRITE"
}

resource "google_compute_disk_attachment" "gcp_storage1_downsized" {
  disk     = google_compute_disk.gcp_storage1_downsized.id
  instance = "instance-1"
  mode     = "READ_WRITE"
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    // Trigger on changes in utilization and cost
    utilization = "14.00"
    cost        = "24.50"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Downsizing gcp_storage1 to 40 GB'"
  }

  connection {
    type        = "ssh"
    host        = "localhost"
    user        = "user"
    private_key = file("~/.ssh/id_rsa")
  }

  depends_on = [google_compute_disk_attachment.gcp_storage1_downsized]
}

resource "google_compute_disk" "gcp_storage1_final" {
  name  = "gcp-storage1-final"
  zone  = "us-central1-a"
  size  = 40 // GB
  type  = "pd-standard"
  labels = {
    env = "dev"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1_final" {
  disk     = google_compute_disk.gcp_storage1_final.id
  instance = "instance-1"
  mode     = "READ_WRITE"
}
```
