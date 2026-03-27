
# Terraform Infra
# Node: aws_vm1
# Time: 1774592364

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

    inline = [
      "sudo yum update -y",
      "sudo yum install -y yum-utils",
      "sudo yum-config-manager --add-repo https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm",
      "sudo yum install -y epel-release",
      "sudo yum install -y amazon-ssm-agent"
    ]
  }
}
