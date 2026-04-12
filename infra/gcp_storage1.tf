
# Terraform Infra
# Node: gcp_storage1
# Time: 1775966780


resource "null_resource" "gcp_storage1" {
  provisioner "local-exec" {
    command = "echo applied DOWNSIZE_MEDIUM"
  }
}

