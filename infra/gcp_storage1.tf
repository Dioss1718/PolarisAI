
# Terraform Infra
# Node: gcp_storage1
# Time: 1774588570

```terraform
resource "google_compute_instance" "gcp_storage1" {
  # ... existing configuration ...

  machine_type = "n1-standard-1"
}

resource "null_resource" "gcp_storage1" {
  triggers = {
    reason  = "Policy Approved"
    action  = "DOWNSIZE_MEDIUM"
    changes = jsonencode({
      utilization = 14.00
      cost        = 24.50
    })
  }
}
```
