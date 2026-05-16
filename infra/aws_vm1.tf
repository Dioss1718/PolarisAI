
# Terraform Infra
# Node: aws_vm1
# Time: 1774801007

resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      host        = "aws_vm1"
      user        = "your_username"
      private_key = file("~/.ssh/your_private_key")
    }

    script = <<-EOF
      sudo yum update -y
      sudo yum install -y yum-utils
      sudo yum-config-manager --enable epel-release
      sudo yum install -y epel-release
      sudo yum update -y
      sudo yum install -y patch
      sudo patch -V
      sudo patch -p0 < /path/to/patch/file.patch
    EOF
  }
}
