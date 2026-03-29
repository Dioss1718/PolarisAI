
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808776

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
    action  = "SECURE_PATCH"
  }

  provisioner "local-exec" {
    command = <<EOF
      gcloud compute instances patch gcp_lb1 --zone=your-zone --maintenance-policy=TERMINATE --min-cpu-metric=cpu-utilization --min-cpu-percent=10 --max-price=63.00 --network-tier=PREMIUM --maintenance-window-start-time=your-time --maintenance-window-duration=your-duration --maintenance-window-recurrence=your-recurrence
    EOF
  }
}

resource "null_resource" "gcp_lb1_exposure" {
  triggers = {
    node_id = "gcp_lb1"
    action  = "SECURE_PATCH"
  }

  provisioner "local-exec" {
    command = <<EOF
      gcloud compute instances patch gcp_lb1 --zone=your-zone --network-tier=PREMIUM --network-interface=0 --network-interface-type=VPC --network-interface-subnet=your-subnet --network-interface-ip-address=your-ip
    EOF
  }
}

resource "null_resource" "gcp_lb1_compliance" {
  triggers = {
    node_id = "gcp_lb1"
    action  = "SECURE_PATCH"
  }

  provisioner "local-exec" {
    command = <<EOF
      gcloud compute instances patch gcp_lb1 --zone=your-zone --labels=your-labels --metadata=your-metadata
    EOF
  }
}
```
