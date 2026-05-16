
# Terraform Infra
# Node: gcp_storage1
# Time: 1774599281

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing properties...

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "gcp_storage1"
    action  = "DOWNSIZE_MEDIUM"
  }
}
```
