# Docker Notes

## Volume Mounts
To create a volume mount:
`docker volume create pgdata`

To list volume mounts:
`docker volume ls`

To inspect the volume mount:
`docker volume inspect pgdata`
Output:
```
[
    {
        "CreatedAt": "2026-07-25T14:10:47Z",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/pgdata/_data",
        "Name": "pgdata",
        "Options": null,
        "Scope": "local"
    }
]
```

_Note_: If you're using Docker Desktop on macOS or Windows, that `/var/lib/docker/...` path does not exist directly on your Mac/Windows filesystem. Docker is actually running inside a small Linux VM, and that mountpoint exists inside Docker's VM.