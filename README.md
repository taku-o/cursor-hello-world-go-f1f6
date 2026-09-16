# cursor-hello-world-go-f1f6

A minimal Go "Hello, World!" HTTP service used to demonstrate a Cursor Cloud Agent development environment.

## Requirements

- Go 1.22+

## Run

```bash
go run .
```

The server listens on port `8080` by default (override with the `PORT` environment variable):

```bash
curl http://localhost:8080/        # -> Hello, World!
curl http://localhost:8080/healthz # -> ok
```

## Test

```bash
go test ./...
```

## Build

```bash
go build ./...
```
