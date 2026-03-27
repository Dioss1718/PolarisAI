
# Terraform Infra
# Node: aws_vm1
# Time: 1774592192

resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      host        = "aws_vm1"
      user        = "ec2-user"
      private_key = file("~/.ssh/aws_vm1_key")
    }

    inline = [
      "sudo yum update -y",
      "sudo yum install -y yum-utils",
      "sudo yum-config-manager --enable AmazonLinux-V2-Packages",
      "sudo yum install -y amazon-linux-extras",
      "sudo amazon-linux-extras install -y aws-cfn-bootstrap",
      "sudo yum clean all",
      "sudo yum install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/AmazonSSMAgentSetup.msi"
    ]
  }
}
