# Yoga Project

Yoga Project is a Go web API for a yoga class booking application. The long-term goal is a full web app where users can create accounts, instructors can publish available yoga class sessions, and users can register for sessions.

This project is also being used as a portfolio learning project, with emphasis on API design, relational database modeling, authentication, middleware, and maintainable backend structure.

## Current Scope

- User account creation
- Instructor account creation
- Instructor-only yoga session creation
- Browser-based login, signup, session calendar, and class registration UI
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
├── api_docs.md
├── yoga_project.http
├── web/
│   ├── index.html
│   ├── styles.css
│   └── app.js
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
INSTRUCTOR_CREATION_TOKEN
ADMIN_KEY
IN_DEV
```

`IN_DEV=true` enables dev-only behavior such as the database reset endpoint.
`FILEPATH_ROOT` can be set to `web`; if omitted, the server defaults to serving frontend files from the `web` directory.

## Database Model

The core schema currently models:

- `users`: application users, including instructors via `is_instructor`
- `sessions`: scheduled yoga classes owned by instructor users
- `class_registrations`: user-to-session registrations
- `refresh_tokens`: refresh tokens tied to users

The current session model uses `sessions.instructor_id -> users.id`, where the referenced user is expected to have `is_instructor = true`.

## API Documentation

See [api_docs.md](api_docs.md) for endpoint details, request bodies, response bodies, and authorization notes.

## Frontend

The frontend is a plain HTML/CSS/JavaScript app in the `web` directory. The Go server serves it from `/`, so the browser and API share the same origin during local development.

Key files:

- `web/index.html`: page structure and form/dialog markup
- `web/styles.css`: visual layout and responsive styling
- `web/app.js`: browser state, API requests, login/signup, calendar rendering, and registration behavior

## Development Notes

Generate database code after changing SQL schema or queries:

```sh
sqlc generate
```

Run tests:

```sh
go test ./...
```

Run the app locally:

```sh
PORT=:8080 FILEPATH_ROOT=web go run .
```

Open the frontend:

```text
http://localhost:8080/
```

## Endpoint Testing

Use [yoga_project.http](yoga_project.http) with a REST Client extension or compatible IDE HTTP runner. It includes requests for:

- resetting the dev database
- creating instructors
- logging in as instructors
- creating sample sessions
- creating/logging in as a student
- listing sessions
- registering and unregistering for a class

Before running it, replace the top-level variables for `adminKey` and `instructorCreationToken` with values matching your local `.env`.

## Design Goals

- Keep domain rules close to the database where constraints are appropriate.
- Keep request handlers readable while extracting reusable concerns over time.
- Use request-aware logging for useful debugging context.
- Avoid storing derived data unless there is a clear performance or consistency reason.
