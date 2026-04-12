
# Terraform Infra
# Node: 
# Time: 1775972239


resource "null_resource" "" {
  provisioner "local-exec" {
    command = "echo applied DOWNSIZE_MEDIUM"
  }
}

