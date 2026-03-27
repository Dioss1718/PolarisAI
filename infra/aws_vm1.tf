
# Terraform Infra
# Node: aws_vm1
# Time: 1774595299

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

    script = <<-EOF
      sudo apt update && sudo apt install -y unattended-upgrades
      sudo unattended-upgrades -y --auto-all
      sudo apt update && sudo apt full-upgrade -y
      sudo apt autoremove -y
      sudo apt autoclean -y
      sudo apt clean -y
      sudo apt update && sudo apt install -y awscli
      sudo aws s3 sync s3://my-bucket/ /var/www/html/
      sudo systemctl restart apache2
    EOF
  }
}
