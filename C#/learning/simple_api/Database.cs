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
}