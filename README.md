# Spy cats

Project structure created based on https://github.com/golang-standards/project-layout.

## How to run

1. `docker compose up` - create container to start postgres

2. `go run cmd/server/*` - run server

## API endpoints

All API endpoints described in [Postman collection](SpyCat.postman_collection.json).

## Possible improvements

1. Add database transaction to create mission with targets logic
2. Add context info to errors (stacktrace or wrapping)
3. Add validation to service args
4. Add request/response logging middleware
5. Cover logic with tests