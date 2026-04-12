
# Terraform Infra
# Node: 
# Time: 1775975593

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
      "sudo yum install -y yum-utils",
      "sudo yum-config-manager --enable rhui-REGION-rhel-server-extras",
      "sudo yum install -y rhui-REGION-rhel-server-extras",
      "sudo yum install -y rhui-REGION-rhel-server-rpms",
      "sudo yum install -y epel-release",
      "sudo yum install -y ansible",
      "sudo yum install -y python3",
      "sudo yum install -y python3-pip",
      "sudo pip3 install --upgrade pip",
      "sudo pip3 install ansible",
      "sudo ansible-galaxy collection install community.general",
      "sudo ansible-playbook -i /etc/ansible/hosts /path/to/playbook.yml",
    ]
  }
}
```
