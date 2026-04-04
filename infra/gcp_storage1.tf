
# Terraform Infra
# Node: gcp_storage1
# Time: 1774808835

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance properties
  // ...

  machine_type = "n1-standard-1"
}

resource "null_resource" "remediation" {
  triggers = {
    node_id = "gcp_storage1"
    action  = "DOWNSIZE_MEDIUM"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --machine-type=n1-standard-1 --zone=${google_compute_instance.gcp_storage1.zone}
EOF
  }
}
```
