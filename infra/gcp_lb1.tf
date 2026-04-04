
# Terraform Infra
# Node: gcp_lb1
# Time: 1774808933

```terraform
resource "null_resource" "gcp_lb1" {
  triggers = {
    node_id = "gcp_lb1"
  }

  provisioner "local-exec" {
    command = <<EOF
gcloud compute instances patch gcp_lb1 --zone us-central1-a --maintenance-policy MIGRATE --min-instances 0 --max-instances 0 --network-tier STANDARD --no-public-ip --no-external-ip --maintenance-window 0:0:0 --redundant-components none --redundant-disks none --redundant-ports none --redundant-interfaces none --redundant-sockets none --redundant-cores none --redundant-memory none --redundant-gpu none --redundant-ssd none --redundant-hdd none --redundant-ssd-count 0 --redundant-hdd-count 0 --redundant-gpu-count 0 --redundant-memory-mb 0 --redundant-cores-count 0 --redundant-ports-count 0 --redundant-interfaces-count 0 --redundant-sockets-count 0 --redundant-gpu-type none --redundant-ssd-type none --redundant-hdd-type none --redundant-ssd-size 0 --redundant-hdd-size 0 --redundant-gpu-size 0 --redundant-memory-size 0 --redundant-cores-size 0 --redundant-ports-size 0 --redundant-interfaces-size 0 --redundant-sockets-size 0 --redundant-gpu-count-per-instance 0 --redundant-ssd-count-per-instance 0 --redundant-hdd-count-per-instance 0 --redundant-memory-mb-per-instance 0 --redundant-cores-count-per-instance 0 --redundant-ports-count-per-instance 0 --redundant-interfaces-count-per-instance 0 --redundant-sockets-count-per-instance 0 --redundant-gpu-type-per-instance none --redundant-ssd-type-per-instance none --redundant-hdd-type-per-instance none --redundant-ssd-size-per-instance 0 --redundant-hdd-size-per-instance 0 --redundant-gpu-size-per-instance 0 --redundant-memory-size-per-instance 0 --redundant-cores-size-per-instance 0 --redundant-ports-size-per-instance 0 --redundant-interfaces-size-per-instance 0 --redundant-sockets-size-per-instance 0 --redundant-gpu-count-per-instance-type 0 --redundant-ssd-count-per-instance-type 0 --redundant-hdd-count-per-instance-type 0 --redundant-memory-mb-per-instance-type 0 --redundant-cores-count-per-instance-type 0 --redundant-ports-count-per-instance-type 0 --redundant-interfaces-count-per-instance-type 0 --redundant-sockets-count-per-instance-type 0 --redundant-gpu-type-per-instance-type none --redundant-ssd-type-per-instance-type none --redundant-hdd-type-per-instance-type none --redundant-ssd-size-per-instance-type 0 --redundant-hdd-size-per-instance-type 0 --redundant-gpu-size-per-instance-type 0 --redundant-memory-size-per-instance-type 0 --redundant-cores-size-per-instance-type
