
# Terraform Infra
# Node: gcp_storage1
# Time: 1775929202

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

  metadata = {
    gcp-storage1-utilization = 14.00
    gcp-storage1-cost        = 24.50
  }

  metadata_startup_script = <<EOF
  gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1 --metadata gcp-storage1-utilization=14.00,gcp-storage1-cost=24.50
  EOF
}
```
