
# Terraform Infra
# Node: gcp_storage1
# Time: 1774600901

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone us-central1-a --machine-type n1-standard-1
EOF
  }
}
```
