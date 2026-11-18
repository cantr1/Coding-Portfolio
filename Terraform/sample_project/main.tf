resource "terraform_data" "example" {
    input = local.greeting
}

locals {
    greeting = "${var.environment}-server"
}