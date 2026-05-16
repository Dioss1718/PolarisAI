
# Terraform Infra
# Node: gcp_lb1
# Time: 1775967305

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      terraform import google_compute_instance gcp_lb1
      terraform state update "resource 'google_compute_instance.gcp_lb1' { ... }" -field 'exposure' -value 'INTERNAL'
      terraform state update "resource 'google_compute_instance.gcp_lb1' { ... }" -field 'cost' -value '63.00'
      terraform state update "resource 'google_compute_instance.gcp_lb1' { ... }" -field 'compliance' -value 'true'
      terraform destroy -target=google_compute_instance.gcp_lb1
    EOF
  }
}

resource "google_compute_instance" "gcp_lb1" {
  name         = "gcp_lb1"
  zone         = "your-zone"
  machine_type = "your-machine-type"

  boot_disk {
    initialize_params {
      image = "your-image"
    }
  }

  network_interface {
    network = "your-network"
  }
}
```
