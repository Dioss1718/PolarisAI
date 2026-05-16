
# Terraform Infra
# Node: gcp_lb1
# Time: 1774801322

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    command = <<EOF
      gcloud compute backend-services update gcp_lb1 --region us-central1 --access-policies ${var.access_policies} --load-balancing-scheme INTERNAL
      gcloud compute backend-services patch gcp_lb1 --region us-central1 --cost 63.00
      gcloud compute backend-services update gcp_lb1 --region us-central1 --compliance ${var.compliance}
    EOF
  }
}
```
