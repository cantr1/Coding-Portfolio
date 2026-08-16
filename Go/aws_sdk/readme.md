# EC2 Instance Management in Go

This repository demonstrates how to manage Amazon EC2 instances using the AWS SDK for Go. It provides functions to list, start, and stop EC2 instances.

This is very useful for me, allowing me to quickly spin up instances when I need them and stop them to save costs.

### Setup
Very simple, before doing anything with this repo, you need to run the following:

`aws configure`

This will prompt you to enter your AWS access key ID, secret access key, default region name, and default output format. Make sure to enter these details correctly.

### Usage
To use this repository, follow these steps:

1. Clone the repository to your local machine.
2. Run `go build` to build the executable.
3. Run the executable and follow the prompts to manage your EC2 instances.

### Tips
I recommend keeping the compiled executable in a directory that is easily accessible, such as your home directory or a dedicated tools directory.

I have an alias in `.bashrc` so that I can simply run `ec2` to execute the compiled executable and manage my servers.