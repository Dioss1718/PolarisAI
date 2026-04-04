
# Terraform Infra
# Node: gcp_storage1
# Time: 1774813131

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1"
  zone  = "us-central1-a"
  size  = 50 // GB
  type  = "pd-standard"
  labels = {
    environment = "production"
  }
}

resource "google_compute_disk" "gcp_storage1_downsized" {
  name  = "gcp-storage-1-downsized"
  zone  = "us-central1-a"
  size  = 40 // GB
  type  = "pd-standard"
  labels = {
    environment = "production"
  }
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    node_id = "gcp_storage1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
      gcloud compute disks resize ${google_compute_disk.gcp_storage1_downsized.name} --size 40 --zone us-central1-a
    EOF
  }
}
```
