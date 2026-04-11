
# Terraform Infra
# Node: azure_vm1
# Time: 1775913789

```terraform
resource "null_resource" "azure_vm1_downsize" {
  triggers = {
    reason  = "Policy Approved"
    node_id = "azure_vm1"
    action   = "DOWNSIZE_MEDIUM"
  }

  connection {
    type        = "ssh"
    host        = "your_azure_vm1_ip"
    user        = "your_azure_vm1_username"
    private_key = file("~/.ssh/your_azure_vm1_private_key")
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      host        = "your_azure_vm1_ip"
      user        = "your_azure_vm1_username"
      private_key = file("~/.ssh/your_azure_vm1_private_key")
    }

    inline = [
      "az vm resize --resource-group your_resource_group --name azure_vm1 --size Standard_DS2_v2",
    ]
  }
}

output "remediation_changes" {
  value = {
    utilization = {
      before  = 15.00
      after   = 10.50
    }
    cost = {
      before  = 195.00
      after   = 136.50
    }
  }
}
```
