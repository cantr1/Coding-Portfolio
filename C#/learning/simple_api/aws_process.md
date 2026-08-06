# Deployment to AWS
This is a general outline of how I deployed this project to AWS, not necessarily in great detail but a rough guide.

### Provisioning
Via the web console, I provisioned a t3.micro EC2 instance in my dev VPC. I setup an SSH key then logged into the server.

I ran the following commands to set it up for Dotnet:
```
sudo dnf update -y

sudo dnf install aspnetcore-runtime-8.0

sudo dnf install -y dotnet-sdk-8.0

dotnet --list-runtimes # verify 

sudo dnf install git
```

### Dotnet Development
I used VSCodes remote development extension to SSH to the server and do my development directly on the machine.

As I started to learn more, I naturally went to loading env vars to handle secrets. I found that you need to 
install the DotNetEnv package from the project folder in order to load variables into the program:
`dotnet add package DotNetEnv`

However, this does not appear to be idiomatic C#. The better way is to let the project read env vars through configuration. This means using the like:
`var port = builder.Configuration["PORT"] ?? "5000";`

Then running the project with:
`PORT=5000 dotnet run`

### Notes
for a development server, installing the SDK is fine. For a production server, you usually publish locally or in CI/CD, deploy the compiled output, and install only the runtime on AWS

This would def be an improvement, but one I am not quite ready for yet