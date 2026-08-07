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
        public User(string name, string username)
        {
            Name = name;
            Username = username;
        }
        public string Name { get; set; }
        public string Username { get; set; }
    }

    [ApiController]
    [Route("/api/health")]
    public class HealthController : ControllerBase
    {
        [HttpGet]
        public ActionResult ReturnHealthStatus()
        {
            return Ok(new { status = "ok" });
        }
    }

    [ApiController]
    [Route("/api/users")]
    public class UsersController : ControllerBase
    {
        private static readonly User DefaultUser = new User("Kelly", "kelz");
        [HttpGet]
        public ActionResult ReturnUsers()
        {
            return Ok(DefaultUser);
        }
    }


}
