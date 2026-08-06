using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Hosting;
using DotNetEnv;

var builder = WebApplication.CreateBuilder(args);
var port = builder.Configuration["PORT"] ?? "5000";
var app = builder.Build();

// Set the port and IP address binding to listen for all requests on the server IP
app.Urls.Add($"http://0.0.0.0:{port}");
app.MapGet("/health", () => Results.Ok(new { status = "ok" }));

app.Run();