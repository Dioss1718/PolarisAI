
# Terraform Infra
# Node: gcp_storage1
# Time: 1775967309

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
  metadata = {
    startup-script = <<-EOF
      #!/bin/bash
      gcloud compute instances update gcp-storage1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}

resource "null_resource" "gcp_storage1" {
  provisioner "local-exec" {
    command = <<EOF
      gcloud compute instances update gcp-storage1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}
```
