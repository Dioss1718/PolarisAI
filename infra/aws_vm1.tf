
# Terraform Infra
# Node: aws_vm1
# Time: 1774598023

resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      host        = "aws_vm1"
      user        = "ubuntu"
      private_key = file("~/.ssh/aws_vm1_key")
    }

    script = <<EOF
      sudo apt update
      sudo apt install -y unattended-upgrades
      sudo unattended-upgrades -y --download-only
      sudo apt dist-upgrade -y
      sudo apt autoremove -y
      sudo apt autoclean -y
      sudo apt clean -y
      sudo apt update
      sudo apt install -y awscli
      sudo aws s3 sync s3://my-bucket/ /tmp/
      sudo aws s3 sync /tmp/ s3://my-bucket/
      sudo apt update
      sudo apt install -y python3-pip
      sudo pip3 install awscli
      sudo aws s3 sync s3://my-bucket/ /tmp/
      sudo aws s3 sync /tmp/ s3://my-bucket/
    EOF
  }
}
