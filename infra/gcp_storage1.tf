
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801952

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

  metadata = {
    startup-script = <<-EOF
      #!/bin/bash
      echo "Instance resized to n1-standard-1"
    EOF
  }
}

resource "null_resource" "resize" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Instance resized to n1-standard-1'"
  }
}
```
