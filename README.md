# User Info CLI

A command-line tool for managing user information. It allows you to create new users and search for existing users by name. The tool also includes an HTTP server that provides a REST API for user operations.

## Features

- Create new users with name and mobile number
- Search for users by name (case-insensitive by default)
- In-memory storage using JSON file
- REST API server for user operations

## Installation

### Prerequisites

- Go 1.25.1 or later

### Build from source

```bash
git clone https://github.com/utpal74/user-info-cli-cobra.git
cd user-info-cli-cobra
go build -o user-info .
```

### Install directly

```bash
go install github.com/utpal74/user-info-cli-cobra@latest
```

## Usage

### CLI Commands

#### Create a user

```bash
./user-info create "John Doe" "1234567890"
```

#### Search for a user

```bash
./user-info search "John"
```

For case-sensitive search:

```bash
./user-info search "John" --case-sensitive
```

#### Get help

```bash
./user-info --help
```

### REST API Server

The project also includes an HTTP server that provides REST endpoints for user operations.

#### Start the server

```bash
go run cmd/main.go
```

The server will start on port 8080.

#### API Endpoints

- `GET /users?name=<name>` - Get user by name
- `POST /users` - Create a new user

Example POST request:

```json
{
  "name": "John Doe",
  "mobile_no": "1234567890"
}
```

## Data Storage

User data is stored in memory and persisted to `users.json` in the following format:

```json
{
  "username": {
    "name": "Full Name",
    "mobile_no": "Phone Number"
  }
}
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router

## License

This project is open source. Please check the license file for details.
