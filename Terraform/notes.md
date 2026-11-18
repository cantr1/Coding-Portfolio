# Terraform Study Notes

# Basics
`terraform init`
Initialize terraform from the current dir

`terraform fmt`
Format and validate the file
`terraform validate`
Validate the configuration

`terraform plan`
View the actions take by terraform before application

`terraform apply`
Apply the plan

`terraform state list`
View resources

`terraform destroy`
Tear down resources

# Providers
`terraform providers`
View current providers in project

`terraform init -upgrade`
Upgrade to highest version allowed

`terraform providers lock -platform=linux_amd64 -platform=linux_arm64`
Lock modules to run on specific platforms