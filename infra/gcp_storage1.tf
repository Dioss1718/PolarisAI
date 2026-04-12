
# Terraform Infra
# Node: gcp_storage1
# Time: 1774805871

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

resource "google_compute_disk_attachment" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = "instance-1234567890"
  mode     = "READ_WRITE"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances update --zone us-central1-a instance-1234567890 --disk 0:size=14GB
EOF
  }
}
```
