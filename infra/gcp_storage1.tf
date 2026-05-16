
# Terraform Infra
# Node: gcp_storage1
# Time: 1774809738

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances downsize --zone us-central1-a --instance gcp-storage1 --machine-type n1-standard-1"
  }
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    utilization = "14.00"
    cost        = "24.50"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Policy Approved: Downsize gcp-storage1 to n1-standard-1 due to utilization ${self.triggers.utilization} and cost ${self.triggers.cost}'"
  }
}
```
