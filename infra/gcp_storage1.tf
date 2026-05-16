
# Terraform Infra
# Node: gcp_storage1
# Time: 1774596441

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 50 // Downsize from 100 to 50 GB
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
}

resource "google_compute_disk" "gcp_storage1_backup" {
  name  = "gcp-storage1-backup"
  zone  = "us-central1-a"
  size  = 100 // Backup size remains the same
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
  depends_on = [google_compute_disk.gcp_storage1]
}

resource "google_compute_disk" "gcp_storage1_snapshot" {
  name  = "gcp-storage1-snapshot"
  zone  = "us-central1-a"
  size  = 100 // Snapshot size remains the same
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
  depends_on = [google_compute_disk.gcp_storage1_backup]
}
```
