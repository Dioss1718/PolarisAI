
# Terraform Infra
# Node: gcp_storage1
# Time: 1774801107

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    instance_size = "n1-standard-1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage1 --zone us-central1-a --machine-type n1-standard-1 --maintenance-policy MIGRATE --min-cpu-platform "Automatic" --boot-disk-size 30 --boot-disk-type pd-standard --boot-disk-device-name gcp-storage1 --reservation-affinity any --labels environment=prod --labels role=storage --labels tier=medium --labels node-id=gcp-storage1 --labels reason=Policy Approved --labels action=DOWNSIZE_MEDIUM --labels utilization=14.00 --labels cost=24.50
    EOF
  }
}
```
