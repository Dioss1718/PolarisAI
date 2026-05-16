
# Terraform Infra
# Node: 
# Time: 1775972233


resource "null_resource" "" {
  provisioner "local-exec" {
    command = "echo applied SECURE_RESTRICT"
  }
}

