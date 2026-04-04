
# Terraform Infra
# Node: gcp_storage1
# Time: 1774813980

```terraform
resource "google_compute_instance" "gcp_storage1" {
  // existing properties...

  machine_type = "n1-standard-1"
}
```

```terraform
resource "google_compute_disk" "gcp_storage1" {
  // existing properties...

  size = 30
}
```

```terraform
resource "google_compute_attached_disk" "gcp_storage1" {
  disk     = google_compute_disk.gcp_storage1.self_link
  instance = google_compute_instance.gcp_storage1.self_link
}
```
