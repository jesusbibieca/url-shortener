# URL shortener written in Go

[![ci-test](https://github.com/jesusbibieca/url-shortener/actions/workflows/ci.yaml/badge.svg)](https://github.com/jesusbibieca/url-shortener/actions/workflows/ci.yaml)

This is a simple URL shortener written in Go. It uses Redis as the DB and the web framework Gin.

## Why?

I'm learning Go and I wanted to build something simple to get started. I want to keep adding features to this project as I learn more about Go and make it a more complete project.

## How to run

**Note:** This is very much a work in progress. I'm still learning Go and I'm sure there are better ways to do things and it will be improving along the way.

### Prerequisites

- Go 1.20+
- Redis
- golang-migrate - This is used to run the migrations. You can install it with

```bash
  brew install golang-migrate
```

- sqlc - This is used to generate the code for the DB. You can install it with

```bash
  brew install sqlc
```

- Reflex (optional) - This is used to auto reload the server when a file changes. You can install it with

```bash
  go install github.com/cespare/reflex
```

### Steps

1. Clone the repo
2. Run `go get` to install dependencies
3. Install and run Redis on port `6379`
4. Copy the `.env.example.yaml` file to `.env.yaml` and update the values `cp .env.example.yaml .env.yaml`
5. Run `make dev` or `make dev-watch` to auto reload on changes
6. Navigate to `localhost:8080` in your browser

## Usage

TODO: Maybe figure out a way to show which endpoints are available and how to use them

## TODO

- [x] Configure auto reload when file changes
- [x] Add better error handling
- [x] Add a logger
- [x] Add a config to set the port and other variables (maybe viper?)
- [x] Add a config to start redis and db with docker-compose (or something similar)
- [x] Add a database to store the URLs long term (maybe Postgres? or just sqlite?)
- [x] handle caching

- [x] Add a way to manage the URLs (CRUD)
  - [x] Create
  - [x] Read
  - [x] Update
  - [x] Delete
  - [x] Expose to API
- [x] Add a way to manage the users (CRUD)

  - [x] Create
  - [x] Read
  - ~~[ ]Update~~
  - [x] Delete
  - [x] Expose to API

- [x] Add authentication
- [-] Add authorization (Implemented on delete url)
  - Missing adding authorization to all other endpoints
- [ ] Only owner deletes/updates URLs

- [ ] Validate the URL before creating it (handle things like `http[s]://google.com` and `google.com`)

- [ ] Add a way to track the number of times a URL is accessed
- [ ] Add a way to track the number of times a URL is accessed by a specific user?
- [ ] Add a way to track the number of times a URL is accessed by a specific user in a specific time period?
- [ ] Log request redirect times

- [ ] Add more tests

- [ ] Add a frontend to create and manage the URLs
  - [ ] Add a way to create a new URL for unknown users
  - [ ] Add a way to create users
  - [ ] Add a way to login
  - [ ] Add a way to logout
  - [ ] Add a way for users to manage their URLs
