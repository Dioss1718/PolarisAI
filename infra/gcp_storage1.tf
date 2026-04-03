
# Terraform Infra
# Node: gcp_storage1
# Time: 1774808461

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance properties...

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone ${google_compute_instance.gcp_storage1.zone} --machine-type n1-standard-1"
  }
}
```
