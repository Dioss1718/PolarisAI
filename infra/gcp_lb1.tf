
# Terraform Infra
# Node: gcp_lb1
# Time: 1775965764

```terraform
resource "google_compute_firewall" "gcp_lb1" {
  name    = "gcp_lb1-firewall"
  network = "default"

  allow {
    protocol = "tcp"
  }

  target_tags = ["gcp_lb1"]
}

resource "google_compute_firewall" "gcp_lb1_internal" {
  name    = "gcp_lb1-internal-firewall"
  network = "default"

  allow {
    protocol = "tcp"
  }

  target_tags = ["gcp_lb1"]
  source_tags = ["gcp_lb1"]
}

resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances patch gcp_lb1 --maintenance-policy=TERMINATE --min-cpu-maintenance-policy=TERMINATE --min-cpu-platform=None --restart-lowest-cpu --zone=us-central1-a --quiet
gcloud compute instances update-guest-attributes gcp_lb1 --zone=us-central1-a --maintenance-policy=TERMINATE --min-cpu-maintenance-policy=TERMINATE --min-cpu-platform=None --restart-lowest-cpu --quiet
gcloud compute instances update gcp_lb1 --zone=us-central1-a --maintenance-policy=TERMINATE --min-cpu-maintenance-policy=TERMINATE --min-cpu-platform=None --restart-lowest-cpu --quiet
EOF
  }
}
```
