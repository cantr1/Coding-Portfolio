# Deployment to AWS
This is a general outline of how I deployed this project to AWS, not necessarily in great detail but a rough guide.

### Provisioning
Via the web console, I provisioned a t3.micro EC2 instance in my dev VPC. I setup an SSH key then logged into the server.

I ran the following commands to set it up for Dotnet:
```
sudo dnf update -y

sudo dnf install aspnetcore-runtime-10.0

dotnet --list-runtimes # verify 

sudo dnf install git
```

### Dotnet Development
I used VSCodes remote development extension to SSH to the server and do my development directly on the machine.