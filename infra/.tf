
# Terraform Infra
# Node: 
# Time: 1775975619

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
    user        = "NODE_USERNAME"
    private_key = file("~/.ssh/NODE_PRIVATE_KEY")
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install -y aws-cfn-bootstrap",
      "sudo /opt/aws/bin/cfn-signal -e 0 --resource ${self.triggers.node_id} --stack ${self.triggers.stack_id}"
    ]
  }
}
```
