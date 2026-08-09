using Npgsql;

namespace simple_api;

public class Database
{
    private readonly string _connectionString;

    public Database(string connectionString)
    {
        this._connectionString = connectionString;
    }

    public async void TestDbConnection()
    {
        try
        {
            // Initialize and automatically dispose of the connection
            await using var conn = new NpgsqlConnection(_connectionString);

            // Open the connection asynchronously
            await conn.OpenAsync();
            Console.WriteLine("Successfully connected to PostgreSQL!");

            // Example query: Fetching data
            string sql = "SELECT version();";
            await using var cmd = new NpgsqlCommand(sql, conn);
            await using var reader = await cmd.ExecuteReaderAsync();

            while (await reader.ReadAsync())
            {
                Console.WriteLine($"DB Version: {reader.GetString(0)}");
            }

        }
        catch (Exception ex)
        {
            Console.WriteLine($"An error occurred: {ex.Message}");
        }
    }

    public async Task<List<User>> GetAllUsers()
    {
        var users = new List<User>();
        try
        {
            // Initialize and automatically dispose of the connection
            await using var conn = new NpgsqlConnection(_connectionString);

            // Open the connection asynchronously
            await conn.OpenAsync();
            Console.WriteLine("Successfully connected to PostgreSQL!");

            // Fetching data
            string sql = "SELECT name, username, email FROM users;";
            await using var cmd = new NpgsqlCommand(sql, conn);
            await using var reader = await cmd.ExecuteReaderAsync();

            while (await reader.ReadAsync())
            {
                // Console.WriteLine($"UUID: {reader.GetGuid(0)}");
                // Console.WriteLine($"User: {reader.GetString(1)}");
                users.Add(new User(reader.GetString(0),
                    reader.GetString(1),
                    reader.GetString(2)));
            }
            return users;

        }
        catch (Exception ex)
        {
            Console.WriteLine($"An error occurred: {ex.Message}");
        }
        return users;
    }

    public async Task<User> CreateUser(string name, string username, string email)
    {
        await using var conn = new NpgsqlConnection(_connectionString);

        // Open the connection asynchronously
        await conn.OpenAsync();

        string sql = @"
        INSERT INTO users (id, name, username, email)
        VALUES (gen_random_uuid(), @name, @username, @email)
        RETURNING name, username, email;
    ";

        await using var cmd = new NpgsqlCommand(sql, conn);
        cmd.Parameters.AddWithValue("name", name);
        cmd.Parameters.AddWithValue("username", username);
        cmd.Parameters.AddWithValue("email", email);
        await using var reader = await cmd.ExecuteReaderAsync();

        // Parse user data from the DB
        if (await reader.ReadAsync())
        {
            return new User(
                reader.GetString(0),
                reader.GetString(1),
                reader.GetString(2)
                );
        }
        throw new Exception("User creation failed");
    }
}