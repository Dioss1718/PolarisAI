
# Terraform Infra
# Node: 
# Time: 1775971099

```terraform
resource "null_resource" "node" {
  triggers = {
    node_id = "NODE_ID"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'SECURE_PATCH: Node ${self.triggers.node_id} remediated due to Policy Approved'"
  }

  connection {
    type        = "ssh"
    host        = "NODE_IP"
    user        = "NODE_USER"
    private_key = file("~/.ssh/NODE_KEY")
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install -y aws-cfn-bootstrap",
      "sudo /opt/aws/bin/cfn-signal -e 0 --resource ${self.triggers.node_id} --region ${aws_region}"
    ]
  }
}
```
