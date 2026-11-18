terraform {
    required_providers {
        local = {
            source = "hashicorp/local"
            version = "~> 2.5"
        }
    }
}

provider "local" {}

resource "local_file" "example" {
    filename = "${path.module}/example.txt"
    content = "Managed by Terraform"
    file_permission = "0644"
}