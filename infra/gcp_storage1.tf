
# Terraform Infra
# Node: gcp_storage1
# Time: 1775928439

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "e2-medium"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage-1 --zone us-central1-a --machine-type e2-small --maintenance-policy MIGRATE --min-cpu-platform "Automatic" --boot-disk-size 30 --boot-disk-type pd-ssd --boot-disk-device-name gcp-storage-1 --reservation-affinity any --labels environment=prod --labels owner=engineering --labels team=storage --labels reason="Policy Approved"
    EOF
  }
}
```
