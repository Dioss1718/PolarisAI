
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801330

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 14 // GB
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    reason = "Policy Approved"
  }
}
```
