# API Development in C#
This is a general outline of how I wrote this API and deployed it to AWS, not necessarily in great detail but a rough guide
and notes on what I learned.

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

### Development Journal
#### Basic Setup
I used VSCode's remote development extension to SSH to the server and do my development directly on the machine.

As I started to learn more, I naturally went to loading env vars to handle secrets. I found that you need to 
install the DotNetEnv package from the project folder in order to load variables into the program:
`dotnet add package DotNetEnv`

However, this does not appear to be idiomatic C#. 
The better way is to let the project read env vars through configuration. This means using the like:
`var port = builder.Configuration["PORT"] ?? "5000";`

Then running the project with:
`PORT=5000 dotnet run`

#### Authentication
The next logical step was to implement authentication for the API with simple tokens.
My go-to for creating keys is to run the following:
`openssl rand -hex 16`

With that, I discovered the use of the singleton pattern. What this means very succintly is that the entire application
shares a single instance of a class. This is useful for managing resources and ensuring that only one instance of a class is created throughout the application's lifetime.

I use that to load the auth token I created and thus allow all of the API endpoints to access the functions
I created to verify the auth token is correct.

For added security, I have a .env file in the root that I use the following command to load in my shell:
`source .env`

With that, I can now run the project and keep the secrets out of the bash history.
`ACCESS_TOKEN=$API_KEY PORT=$PORT dotnet run`

### Notes
for a development server, installing the SDK is fine. For a production server, you usually publish locally or in CI/CD, deploy the compiled output, and install only the runtime on AWS

This would def be an improvement, but one I am not quite ready for yet