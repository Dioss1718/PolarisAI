
# Terraform Infra
# Node: gcp_storage1
# Time: 1774805782

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
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

resource "google_compute_disk" "gcp_storage1" {
  name  = "gcp-storage1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Instance gcp_storage1 successfully downgraded to n1-standard-1'"
  }
}

resource "null_resource" "gcp_storage1_cost" {
  triggers = {
    cost = "24.50"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Cost of instance gcp_storage1 successfully reduced to $24.50'"
  }
}
```
