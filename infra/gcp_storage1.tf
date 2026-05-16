
# Terraform Infra
# Node: gcp_storage1
# Time: 1774804936

```terraform
resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    node_id = "gcp_storage1"
    reason  = "Policy Approved"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp-storage-1 --machine-type=n1-standard-1 --zone=your-zone
    EOF
  }
}
```
