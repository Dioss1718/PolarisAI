
# Terraform Infra
# Node: manual_vm1
# Time: 1774600745

```terraform
resource "null_resource" "manual_vm1" {
  triggers = {
    node_id = "manual_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      terraform import aws_instance.manual_vm1 manual_vm1
      terraform taint aws_instance.manual_vm1
      terraform apply -auto-approve
      terraform import aws_instance.manual_vm1 manual_vm1
      terraform apply -auto-approve -var 'instance_type=t3.micro'
    EOF
  }
}
```
