
# Terraform Infra
# Node: gcp_lb1
# Time: 1775927164


resource "null_resource" "gcp_lb1" {
  provisioner "local-exec" {
    command = "echo applied SECURE_PATCH"
  }
}

