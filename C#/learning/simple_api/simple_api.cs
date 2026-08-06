using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Mvc;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

var port = builder.Configuration["PORT"] ?? "5000";
var app = builder.Build();

// Set the port and IP address binding to listen for all requests on the server IP.
app.Urls.Add($"http://0.0.0.0:{port}");

app.MapControllers();

app.Run();

namespace TestWebApi
{
    public class User
    {
        public required string Name { get; set; }
        public required string Username { get; set; }
    }

    [ApiController]
    [Route("/api/health")]
    public class UsersController : ControllerBase
    {
        [HttpGet]
        public ActionResult ReturnHealthStatus()
        {
            return Ok(new { status = "ok"});
        }
    }
}
