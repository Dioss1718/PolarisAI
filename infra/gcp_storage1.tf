
# Terraform Infra
# Node: gcp_storage1
# Time: 1774808523

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1"
  zone  = "us-central1-a"
  size  = 50 // Downsize from 100 to 50 GB
  type  = "pd-standard"
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = "instance-1"
  mode     = "READ_WRITE"
}
```
