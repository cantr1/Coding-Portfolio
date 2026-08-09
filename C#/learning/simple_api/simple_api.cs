using Microsoft.AspNetCore.Mvc;
using simple_api;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

var port = builder.Configuration["PORT"] ?? "5000";
var accessToken = builder.Configuration["ACCESS_TOKEN"];
if (string.IsNullOrEmpty(accessToken))
{
    throw new Exception("ACCESS_TOKEN is not set in the configuration.");
}
builder.Services.AddSingleton(new TokenValidator(accessToken));
var app = builder.Build();

// Set the port and IP address binding to listen for all requests on the server IP.
app.Urls.Add($"http://localhost:{port}");

app.MapControllers();

app.Run();

namespace TestWebApi
{
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
        private readonly TokenValidator _tokenValidator;

        public UsersController(TokenValidator tokenValidator)
        {
            _tokenValidator = tokenValidator;
        }
        
        // Track users
        private static readonly List<User> Users = new List<User>();
        
        // Return all users
        [HttpGet]
        public ActionResult ReturnUsers()
        {
            if (!_tokenValidator.IsValid(Request))
            {
                return Unauthorized();
            }
            return Ok(Users);
        }
        
        // Create a new user
        [HttpPost]
        //Bind the JSON body to the C# object using [FromBody]
        public IActionResult CreateUser([FromBody] User request)
        {
            if (!_tokenValidator.IsValid(Request))
            {
                return Unauthorized();
            }
            // Validate expected data
            if (string.IsNullOrEmpty(request.Name) || 
                string.IsNullOrEmpty(request.Username) || 
                string.IsNullOrEmpty(request.Email))
            {
                return BadRequest("Invalid user data");
            }
            var tmpUser = new User(request.Name, request.Username, request.Email);
            Users.Add(tmpUser);
            return Ok(tmpUser);
        }
    }


}
