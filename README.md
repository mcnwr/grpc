# Go gRPC Message Service

A simple and efficient gRPC service for managing messages, built with Go. This service provides endpoints to submit and retrieve messages with proper error handling and logging.

## Features

- gRPC server with HTTP/2 support
- Protocol Buffer message definitions
- Thread-safe message storage
- Comprehensive error handling
- Request logging
- gRPC reflection support

## Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (protoc)
- gRPC tools
- Postman (for testing) or grpcurl

## Installation

1. Install Protocol Buffers compiler:

```bash
# For Windows (using chocolatey)
choco install protoc

# For macOS
brew install protobuf

# For Linux
apt-get install protobuf-compiler
```

2. Install Go plugins for protoc:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. Generate gRPC code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/message.proto
```

4. Run the server:

```bash
go run main.go
```

The server will start on `localhost:8080`

## API Documentation

### 1. Submit Message

Submit a new message to the service.

**Method:** `SubmitMessage`

**Request:**

```protobuf
message SubmitMessageRequest {
  string name = 1;
  string message = 2;
}
```

**Response:**

```protobuf
message SubmitMessageResponse {
  bool success = 1;
  string message = 2;
}
```

### 2. Get Message

Retrieve a message by name.

**Method:** `GetMessage`

**Request:**

```protobuf
message GetMessageRequest {
  string name = 1;
}
```

**Response:**

```protobuf
message GetMessageResponse {
  string name = 1;
  string message = 2;
}
```

## Testing with Postman

1. Open Postman
2. Create a new request
3. Set the request type to gRPC
4. Enter the server address: `localhost:8080`
5. Select the service: `MessageService`
6. Choose the method you want to test

### Example: Get Message

1. Select `GetMessage` method
2. Set the request body:

```json
{
  "name": "John"
}
```

### Example: Submit Message

1. Select `SubmitMessage` method
2. Set the request body:

```json
{
  "name": "John",
  "message": "Hello, World!"
}
```

## Testing with grpcurl

### List Services

```bash
grpcurl -plaintext localhost:8080 list
```

### Get Message

```bash
grpcurl -plaintext -d '{"name": "John"}' localhost:8080 api.MessageService/GetMessage
```

### Submit Message

```bash
grpcurl -plaintext -d '{"name": "John", "message": "Hello, World!"}' localhost:8080 api.MessageService/SubmitMessage
```

## Error Handling

The service includes comprehensive error handling for various scenarios:

1. **Invalid Input**: Returns INVALID_ARGUMENT (3)
2. **Message Not Found**: Returns NOT_FOUND (5)
3. **Internal Errors**: Returns INTERNAL (13)

## Project Structure

```
.
├── api/
│   └── message.proto    # Protocol Buffer definitions
├── main.go             # Main server implementation
├── README.md           # This documentation
└── go.mod             # Go module file
```

## Contributing

Feel free to submit issues and enhancement requests!
