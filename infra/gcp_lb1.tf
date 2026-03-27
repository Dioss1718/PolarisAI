
# Terraform Infra
# Node: gcp_lb1
# Time: 1774595295

```terraform
resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp-lb1"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }

  // Apply the SECURE_PATCH action
  lifecycle {
    ignore_changes = [
      // Ignore changes to exposure and cost
      exposure,
      cost,
      // Ignore compliance updates
      compliance
    ]
  }

  // Update exposure to INTERNAL
  exposure = "INTERNAL"

  // Update cost to 63.00
  cost = 63.00

  // Update compliance
  compliance = "policy-approved"
}
```
