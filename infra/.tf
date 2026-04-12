
# Terraform Infra
# Node: 
# Time: 1775972291

```terraform
resource "null_resource" "node_restriction" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    command = <<EOF
      aws ec2 modify-instance-attribute --instance-id NODE_ID --attribute instanceRootDeviceModifier --value /dev/sda1
      aws ec2 modify-instance-attribute --instance-id NODE_ID --attribute instanceBlockDeviceMapping --modify BlockDeviceMapping {
        DeviceName = "/dev/sda1"
        Ebs {
          DeleteOnTermination = true
        }
      }
    EOF
  }
}
```
