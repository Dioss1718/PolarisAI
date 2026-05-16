
# Terraform Infra
# Node: gcp_lb1
# Time: 1774801003

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute backend-services update gcp_lb1 --region <REGION> --security-policy SECURE_POLICY --security-scanner EXPOSED_SCANNER --exposure INTERNAL --cost 63.00 --description "Policy Approved"
EOF
  }
}
```
