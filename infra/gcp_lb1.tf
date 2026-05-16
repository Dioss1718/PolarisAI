
# Terraform Infra
# Node: gcp_lb1
# Time: 1774809246

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    when    = "destroy"
    command = <<EOF
      gcloud compute backend-services update gcp_lb1 --region=REGION --security-policy=SECURE_POLICY
      gcloud compute backend-services update gcp_lb1 --region=REGION --access-policies=ACCESS_POLICY
      gcloud compute backend-services update gcp_lb1 --region=REGION --quota-cost=63.00
    EOF
  }

  provisioner "local-exec" {
    when    = "create"
    command = <<EOF
      gcloud compute backend-services update gcp_lb1 --region=REGION --security-policy=SECURE_POLICY
      gcloud compute backend-services update gcp_lb1 --region=REGION --access-policies=ACCESS_POLICY
      gcloud compute backend-services update gcp_lb1 --region=REGION --access=INTERNAL
      gcloud compute backend-services update gcp_lb1 --region=REGION --quota-cost=63.00
      gcloud compute backend-services update gcp_lb1 --region=REGION --update-labels=COMPLIANCE_APPROVED
    EOF
  }
}
```
