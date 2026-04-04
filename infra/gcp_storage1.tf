
# Terraform Infra
# Node: gcp_storage1
# Time: 1774808939

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance configuration
}

resource "null_resource" "remediation" {
  triggers = {
    node_id = "gcp_storage1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone=${google_compute_instance.gcp_storage1.zone} --machine-type=n1-standard-1
    EOF
  }
}
```
