
# Terraform Infra
# Node: gcp_lb1
# Time: 1774805774

```terraform
resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp-lb1"
  machine_type = "f1-micro"

  // ... existing properties ...

  network_interface {
    network = "default"

    access_config {
      // No change in access config
    }
  }

  // Update exposure to INTERNAL
  network_interface {
    network = "default"

    access_config {
      access_type = "INTERNAL"
    }
  }

  // Update cost to 63.00
  metadata = {
    cost = 63.00
  }
}

resource "null_resource" "gcp_lb1" {
  triggers = {
    reason = "Policy Approved"
  }
}
```
