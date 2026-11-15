resource "terraform_data" "lab_test" {
  input = "Terraform lab is ready"
}

output "result" {
  value = terraform_data.lab_test.input
}

output "env_message" {
  value = var.environment
}