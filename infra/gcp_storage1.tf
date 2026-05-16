
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801638

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 50 // Downsize from 100 GB to 50 GB
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
}
