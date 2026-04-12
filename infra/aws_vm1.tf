
# Terraform Infra
# Node: aws_vm1
# Time: 1775966271

```terraform
resource "null_resource" "aws_vm1" {
  triggers = {
    node_id = "aws_vm1"
  }

  provisioner "local-exec" {
    command = <<EOF
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceRootDeviceModifier --value "/dev/sda1"
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceBlockDeviceMapping --block-device-mappings "[{\"DeviceName\":\"/dev/sda1\",\"Ebs\":{\"DeleteOnTermination\":true}}]"
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceBlockDeviceMapping --block-device-mappings "[{\"DeviceName\":\"/dev/sda1\",\"Ebs\":{\"DeleteOnTermination\":true,\"VolumeSize\":30}}]"
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceBlockDeviceMapping --block-device-mappings "[{\"DeviceName\":\"/dev/sda1\",\"Ebs\":{\"DeleteOnTermination\":true,\"VolumeSize\":30,\"Iops\":300}}]"
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceBlockDeviceMapping --block-device-mappings "[{\"DeviceName\":\"/dev/sda1\",\"Ebs\":{\"DeleteOnTermination\":true,\"VolumeSize\":30,\"Iops\":300,\"Throughput\":125}}]"
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceBlockDeviceMapping --block-device-mappings "[{\"DeviceName\":\"/dev/sda1\",\"Ebs\":{\"DeleteOnTermination\":true,\"VolumeSize\":30,\"Iops\":300,\"Throughput\":125,\"VolumeType\":\"gp2\"}}]"
      aws ec2 modify-instance-attribute --instance-id <instance_id> --attribute instanceBlockDeviceMapping --block-device-mappings "[{\"DeviceName\":\"/dev/sda1\",\"Ebs\":{\"DeleteOnTermination\":true,\"VolumeSize\":30,\"Iops\":300,\"Throughput\":125,\"VolumeType\":\"gp2\",\"Encrypted\":true}}]"
    EOF
    environment = {
      AWS_ACCESS_KEY_ID = "<access_key>"
      AWS_SECRET_ACCESS_KEY = "<secret_key>"
    }
  }
}
```
