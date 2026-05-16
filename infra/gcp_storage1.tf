
# Terraform Infra
# Node: gcp_storage1
# Time: 1774601260

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance configuration
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_id = google_compute_instance.gcp_storage1.id
  }

  provisioner "local-exec" {
    when    = destroy
    command = "gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone ${google_compute_instance.gcp_storage1.zone} --machine-type n1-standard-4"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Policy Approved: Downsize from 20.00 to 14.00 utilization and from 35.00 to 24.50 cost' > remediation.log"
  }
}
```
