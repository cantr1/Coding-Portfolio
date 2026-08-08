namespace simple_api;

public class User
{
    public User(string name, string username, string email)
    {
        Name = name;
        Username = username;
        Email = email;
    }
    public string Name { get; set; }
    public string Username { get; set; }
    public string Email { get; set; }
}