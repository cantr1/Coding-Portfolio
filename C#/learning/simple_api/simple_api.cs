using Microsoft.AspNetCore.Mvc;
using simple_api;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

// Get configured variables
var port = builder.Configuration["PORT"] ?? "5000";
var accessToken = builder.Configuration["ACCESS_TOKEN"];
if (string.IsNullOrEmpty(accessToken))
{
    throw new Exception("ACCESS_TOKEN is not set in the configuration.");
}
builder.Services.AddSingleton(new TokenValidator(accessToken));

// Set up database connection string
var dbHost = builder.Configuration["DB_HOST"];
var dbUsername = builder.Configuration["DB_USERNAME"] ?? "postgres";
var dbPassword = builder.Configuration["DB_PASSWORD"];
var dbName = builder.Configuration["DB_DATABASE"];
var dbPort = builder.Configuration["DB_PORT"] ?? "5432";
var connString = $"Host={dbHost};Port={dbPort};Username={dbUsername};Password={dbPassword};Database={dbName}";
builder.Services.AddSingleton(new Database(connString));

var app = builder.Build();

// Set the port and IP address binding to listen for all requests on the server IP.
app.Urls.Add($"http://0.0.0.0:{port}");

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
    [Route("/api/db_health")]
    public class DbHealthController : ControllerBase
    {
        private readonly TokenValidator _tokenValidator;
        private readonly Database _db;

        public DbHealthController(TokenValidator tokenValidator, Database db)
        {
            _tokenValidator = tokenValidator;
            _db = db;
        }

        [HttpGet]
        public ActionResult ReturnDbHealthStatus()
        {
            if (!_tokenValidator.IsValid(Request))
            {
                return Unauthorized();
            }
            _db.TestDbConnection();
            return Ok(new { status = "ok" });
        }
    }

    [ApiController]
    [Route("/api/users")]
    public class UsersController : ControllerBase
    {
        private readonly TokenValidator _tokenValidator;
        private readonly Database _db;

        public UsersController(TokenValidator tokenValidator, Database db)
        {
            _tokenValidator = tokenValidator;
            _db = db;
        }

        // Return all users
        [HttpGet]
        public async Task<ActionResult<List<User>>> ReturnUsers()
        {
            if (!_tokenValidator.IsValid(Request))
            {
                return Unauthorized();
            }
            List<User> userDb = await _db.GetAllUsers();
            return Ok(userDb);
        }

        // Create a new user
        [HttpPost]
        //Bind the JSON body to the C# object using [FromBody]
        public async Task<IActionResult> CreateUser([FromBody] User request)
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
            User userDb = await _db.CreateUser(request.Name, request.Username, request.Email);
            return Ok(userDb);
        }
    }


}
