using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Hosting;
using DotNetEnv;

Env.Load();

// Read vars
string serverIP = Environment.GetEnvironmentVariable("server_ip");
string serverPort = Environment.GetEnvironmentVariable("server_port");

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

// Set the port and IP address binding to listen for all requests on the server IP
app.Urls.Add($"http://{server_ip}:{server_port}");
app.MapGet("/health", () => Results.Ok(new { status = "ok" }));

app.Run();