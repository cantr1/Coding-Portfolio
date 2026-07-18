# API Docs

Base URL during local development depends on `PORT`.

All request and response bodies are JSON unless otherwise noted.

## Authentication

The API currently uses bearer tokens in the `Authorization` header:

```text
Authorization: Bearer <token>
```

Some endpoints use setup/admin tokens from environment variables. Session creation uses a user JWT and checks that the authenticated user is an instructor.

## Health

### `GET /api/health`

Returns API health status.

#### Response `200`

```json
{
  "status": "online"
}
```

## Metrics

### `GET /api/metrics`

Returns API usage counters.

Requires `ADMIN_KEY` as a bearer token.

#### Response `200`

```json
{
  "file_server_hits": 0,
  "user_creation_hits": 0,
  "instructor_creation_hits": 0,
  "session_creation_hits": 0,
  "class_registration_hits": 0
}
```

#### Errors

- `401 Unauthorized`: missing or invalid bearer token
- `403 Forbidden`: bearer token does not match `ADMIN_KEY`

## Users

### `POST /api/users`

Creates a standard user.

Requires `USER_CREATION_TOKEN` as a bearer token.

#### Request Body

```json
{
  "email": "student@example.com",
  "password": "example-password"
}
```

#### Response `201`

```json
{
  "user_id": "00000000-0000-0000-0000-000000000000",
  "email": "student@example.com",
  "created_at": "2026-07-18T12:00:00Z",
  "updated_at": "2026-07-18T12:00:00Z"
}
```

#### Errors

- `400 Bad Request`: invalid JSON or missing required fields
- `401 Unauthorized`: missing bearer token
- `403 Forbidden`: bearer token does not match `USER_CREATION_TOKEN`
- `500 Internal Server Error`: password hashing, database, or response encoding failure

## Instructors

### `POST /api/instructors`

Creates an instructor account.

Requires `INSTRUCTOR_CREATION_TOKEN` as a bearer token.

#### Request Body

```json
{
  "email": "teacher@example.com",
  "password": "example-password",
  "name": "Teacher Name"
}
```

#### Response `201`

```json
{
  "user_id": "00000000-0000-0000-0000-000000000000",
  "email": "teacher@example.com",
  "instructor_name": "Teacher Name",
  "created_at": "2026-07-18T12:00:00Z",
  "updated_at": "2026-07-18T12:00:00Z"
}
```

#### Errors

- `400 Bad Request`: invalid JSON or missing required fields
- `401 Unauthorized`: missing bearer token
- `403 Forbidden`: bearer token does not match `INSTRUCTOR_CREATION_TOKEN`
- `500 Internal Server Error`: password hashing, database, or response encoding failure

## Sessions

### `POST /api/sessions`

Creates a yoga class session.

Requires a valid user JWT as a bearer token. The authenticated user must have `is_instructor = true`.

#### Request Body

```json
{
  "start_time": "2026-07-18T14:00:00Z",
  "end_time": "2026-07-18T15:00:00Z",
  "difficulty": 3,
  "class_size": 12,
  "description": "Vinyasa flow focused on balance and breath."
}
```

#### Response `201`

```json
{
  "id": "00000000-0000-0000-0000-000000000000",
  "created_at": "2026-07-18T12:00:00Z",
  "updated_at": "2026-07-18T12:00:00Z",
  "start_time": "2026-07-18T14:00:00Z",
  "end_time": "2026-07-18T15:00:00Z",
  "instructor_id": "00000000-0000-0000-0000-000000000000",
  "difficulty": 3,
  "class_size": 12,
  "description": "Vinyasa flow focused on balance and breath."
}
```

#### Validation

- `start_time` is required.
- `end_time` is required and must be after `start_time`.
- `difficulty` must be between `1` and `5`.
- `class_size` must be greater than `0`.
- `description` is required.

#### Errors

- `400 Bad Request`: invalid JSON or validation failure
- `401 Unauthorized`: missing or invalid JWT
- `403 Forbidden`: authenticated user is not an instructor
- `404 Not Found`: token subject does not match a user
- `500 Internal Server Error`: database or response encoding failure

## Dev Reset

### `POST /api/reset`

Deletes development data and resets API metrics.

Requires:

- `IN_DEV=true`
- `ADMIN_KEY` as a bearer token

#### Response `204`

No response body.

#### Reset Order

The endpoint deletes dependent records before parent records:

```text
class_registrations
sessions
users
```

#### Errors

- `401 Unauthorized`: missing bearer token
- `403 Forbidden`: not in dev mode or invalid admin key
- `500 Internal Server Error`: database reset failure

## Planned Endpoints

These are likely next additions as the project grows:

- `POST /api/login`
- `POST /api/refresh`
- `POST /api/logout`
- `GET /api/sessions`
- `GET /api/sessions/{id}`
- `POST /api/sessions/{id}/registrations`
- `DELETE /api/sessions/{id}/registrations`
