
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801127

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance configuration
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<-EOT
      gcloud compute instances update ${google_compute_instance.gcp_storage1.name} \
        --zone ${google_compute_instance.gcp_storage1.zone} \
        --machine-type f1-micro
    EOT
  }
}
```
