
# Terraform Infra
# Node: gcp_storage1
# Time: 1774808676

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage-1"
  zone  = "us-central1-a"
  size  = 50 // Changed from 100 to 50 GB
  type  = "pd-standard"
  labels = {
    node_id = "gcp_storage1"
  }
}

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = "instance-1"
  mode     = "READ_WRITE"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update instance-1 --zone us-central1-a --disk size=50
    EOF
  }
}
```
