
# Terraform Infra
# Node: gcp_storage1
# Time: 1774592202

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1
EOF
  }
}
```
