
# Terraform Infra
# Node: gcp_storage1
# Time: 1774809148

```terraform
resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1"
  zone  = "us-central1-a"
  size  = 50 // GB
  type  = "pd-standard"
  labels = {
    environment = "production"
  }
}

resource "google_compute_disk" "gcp_storage1_downsized" {
  name  = "gcp-storage1-downsized"
  zone  = "us-central1-a"
  size  = 40 // GB
  type  = "pd-standard"
  labels = {
    environment = "production"
  }
}

resource "google_compute_disk" "gcp_storage1_final" {
  name  = "gcp-storage1-final"
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
gcloud compute disks resize --size 40 --zone us-central1-a gcp-storage1-downsized
gcloud compute disks update --size 40 --zone us-central1-a gcp-storage1-final
gcloud compute disks delete --zone us-central1-a gcp-storage1-downsized
EOF
  }
}
```
