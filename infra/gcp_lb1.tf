
# Terraform Infra
# Node: gcp_lb1
# Time: 1775928434

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
    action  = "SECURE_PATCH"
  }

  provisioner "local-exec" {
    command = <<EOF
      gcloud compute instances patch gcp_lb1 --zone=your-zone --maintenance-policy=TERMINATE --boot-disk-size=63.00GB --network-interface=network-tier=PREEMPTIVE --network-interface=access-configs=0 --network-interface=access-configs.0.type=ONE_TO_ONE_NAT --network-interface=access-configs.0.name=your-access-config
    EOF
  }
}
```
