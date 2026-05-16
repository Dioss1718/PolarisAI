
# Terraform Infra
# Node: azure_vm1
# Time: 1775967294

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    node_id = "azure_vm1"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "echo 'Downsizing azure_vm1 from medium to small'"
  }

  connection {
    host        = "localhost"
    type        = "ssh"
    user        = "your_username"
    private_key = file("~/.ssh/your_private_key")
  }

  provisioner "remote-exec" {
    connection = self.connection
    inline = [
      "az vm resize --resource-group your_resource_group --name azure_vm1 --size Standard_DS2_v2",
    ]
  }
}
```
