
# Terraform Infra
# Node: gcp_storage1
# Time: 1774802255

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing instance properties...

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "gcp_storage1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances update gcp_storage1 --zone ${google_compute_instance.gcp_storage1.zone} --machine-type n1-standard-1
    EOF
  }
}
```
