config: dict = {
    "database": {
        "host": "localhost",
        "credentials": {
            "username": "admin",
            "password": "secret"
        }
    },
    "logging": {
        "level": "INFO"
    }
}

def find_key(data: dict, target_key: str) -> any | None:
    print(data)
    for key, value in data.items():
        if key == target_key:
            return value
        else:
            if isinstance(value, dict):
                result = find_key(value, target_key)
                if result is not None:
                    return result
                else:
                    continue
    return None

print(find_key(config, "username"))
print(find_key(config, "host"))
print(find_key(config, "level"))
print(find_key(config, "missing"))