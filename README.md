# Climber Info

A gRPC service in Go that returns climber information. The server listens on port `8080` and provides a `GetClimber` RPC that returns a fun fact about a given climber name.

## Project Structure

```
.
├── client/          # gRPC client
├── protobuf/        # Protocol Buffer definitions and generated code
├── server/          # gRPC server
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.26+
- `protoc` with Go plugins (`protoc-gen-go`, `protoc-gen-go-grpc`)

## Makefile Commands

| Command | Description |
|---|---|
| `make build` | Build both server and client binaries into `bin/` |
| `make build-server` | Build the server binary only |
| `make build-client` | Build the client binary only |
| `make run-server` | Run the gRPC server |
| `make run-client` | Run the gRPC client (requires running server) |
| `make proto` | Regenerate Go code from `.proto` files |
| `make clean` | Remove the `bin/` directory |
