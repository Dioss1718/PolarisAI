
# Terraform Infra
# Node: gcp_storage1
# Time: 1774600082

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance properties...

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    reason  = "Policy Approved"
    action  = "DOWNSIZE_MEDIUM"
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances resize ${google_compute_instance.gcp_storage1.name} --zone=${google_compute_instance.gcp_storage1.zone} --machine-type=n1-standard-1 --maintenance-policy=MIGRATE --disk 0 --disk-size=30 --quiet
EOF
  }
}
```
