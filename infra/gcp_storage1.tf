
# Terraform Infra
# Node: gcp_storage1
# Time: 1774600706

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance resource
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update ${google_compute_instance.gcp_storage1.name} \
        --zone ${google_compute_instance.gcp_storage1.zone} \
        --machine-type n1-standard-4
    EOF
  }
}
```
