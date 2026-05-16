
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801230

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance properties...

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage1 --machine-type=n1-standard-1 --zone=${google_compute_instance.gcp_storage1.zone}
    EOF
  }
}
```
