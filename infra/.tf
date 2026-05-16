
# Terraform Infra
# Node: 
# Time: 1775971056


resource "null_resource" "" {
  provisioner "local-exec" {
    command = "echo applied DOWNSIZE_SMALL"
  }
}

