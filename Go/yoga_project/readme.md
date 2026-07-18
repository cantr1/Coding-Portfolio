# Yoga Project

Yoga Project is a Go web API for a yoga class booking application. The long-term goal is a full web app where users can create accounts, instructors can publish available yoga class sessions, and users can register for sessions.

This project is also being used as a portfolio learning project, with emphasis on API design, relational database modeling, authentication, middleware, and maintainable backend structure.

## Current Scope

- User account creation
- Instructor account creation
- Instructor-only yoga session creation
- API health check
- Basic API usage metrics
- Dev-only database reset endpoint
- PostgreSQL schema managed with Goose-style migrations
- SQL access generated with sqlc

## Tech Stack

- Go
- Standard library `net/http`
- PostgreSQL
- sqlc
- Goose-style SQL migrations
- Argon2id password hashing
- JWT and refresh-token auth foundations

## Project Structure

```text
.
├── api.go
├── internal/
│   ├── auth.go
│   └── database/
├── sql/
│   ├── queries/
│   └── schema/
├── sqlc.yaml
├── go.mod
└── go.sum
```

## Environment Variables

The API currently expects these environment variables:

```text
PORT
DB_URL
TOKEN_DURATION
TOKEN_SECRET
USER_CREATION_TOKEN
INSTRUCTOR_CREATION_TOKEN
ADMIN_KEY
IN_DEV
```

`IN_DEV=true` enables dev-only behavior such as the database reset endpoint.

## Database Model

The core schema currently models:

- `users`: application users, including instructors via `is_instructor`
- `sessions`: scheduled yoga classes owned by instructor users
- `class_registrations`: user-to-session registrations
- `refresh_tokens`: refresh tokens tied to users

The current session model uses `sessions.instructor_id -> users.id`, where the referenced user is expected to have `is_instructor = true`.

## API Documentation

See [api_docs.md](api_docs.md) for endpoint details, request bodies, response bodies, and authorization notes.

## Development Notes

Generate database code after changing SQL schema or queries:

```sh
sqlc generate
```

Run tests:

```sh
go test ./...
```

## Design Goals

- Keep domain rules close to the database where constraints are appropriate.
- Keep request handlers readable while extracting reusable concerns over time.
- Use request-aware logging for useful debugging context.
- Avoid storing derived data unless there is a clear performance or consistency reason.
