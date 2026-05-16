
# Terraform Infra
# Node: gcp_storage1
# Time: 1774818890

resource "google_compute_instance" "gcp_storage1" {
  name         = "gcp-storage-1"
  machine_type = "n1-standard-2"
  zone         = "us-central1-a"
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }
  network_interface {
    network = "default"
  }
}

resource "google_compute_disk" "gcp_storage1_disk" {
  name  = "gcp-storage-1-disk"
  zone  = "us-central1-a"
  size  = 30
  type  = "pd-standard"
}

resource "google_compute_attached_disk" "gcp_storage1_disk_attach" {
  disk     = google_compute_disk.gcp_storage1_disk.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}

resource "null_resource" "gcp_storage1_downsize" {
  triggers = {
    instance_size = "n1-standard-2"
  }

  provisioner "local-exec" {
    when    = destroy
    command = <<EOF
      gcloud compute instances resize gcp-storage-1 --zone us-central1-a --machine-type n1-standard-1
    EOF
  }
}
