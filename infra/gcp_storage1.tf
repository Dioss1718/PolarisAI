
# Terraform Infra
# Node: gcp_storage1
# Time: 1774609443

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage1"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    reason = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp-storage1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}
```
