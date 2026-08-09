namespace simple_api;

public sealed class TokenValidator
{
    private readonly string _expectedToken;

    public TokenValidator(string expectedToken)
    {
        _expectedToken = expectedToken;
    }

    public bool IsValid(HttpRequest request)
    {
        var authorization = request.Headers.Authorization.ToString();
        const string bearerPrefix = "Bearer ";

        return authorization.StartsWith(bearerPrefix, StringComparison.OrdinalIgnoreCase)
            && string.Equals(
                authorization[bearerPrefix.Length..],
                _expectedToken,
                StringComparison.Ordinal);
    }
}
