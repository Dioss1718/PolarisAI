
# Terraform Infra
# Node: gcp_storage1
# Time: 1775913801

```terraform
resource "google_compute_disk" "gcp_storage1" {
  zone      = "your-zone"
  name      = "gcp-storage1"
  size      = 50 // in GB
  zone      = "your-zone"
  disk_size = 50 // in GB
}

resource "google_compute_disk" "gcp_storage1_downsized" {
  zone      = "your-zone"
  name      = "gcp-storage1-downsized"
  size      = 25 // in GB
  zone      = "your-zone"
  disk_size = 25 // in GB
}

resource "google_compute_disk" "gcp_storage1_final" {
  zone      = "your-zone"
  name      = "gcp-storage1-final"
  size      = 25 // in GB
  zone      = "your-zone"
  disk_size = 25 // in GB
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    node_id = "gcp_storage1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute disks resize --zone your-zone gcp-storage1 --size 25 --quiet
gcloud compute disks update --zone your-zone gcp-storage1 --labels=utilization=14.00,cost=24.50
EOF
  }
}
```
