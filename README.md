# gRPC Web Interface

A web-based interface for testing gRPC connections, allowing users to submit and retrieve messages through a user-friendly HTML interface.

## Project Structure

```
.
├── api/
│   └── message.proto      # Protocol buffer definitions
├── templates/
│   └── grpc_test.html    # Web interface template
├── web/
│   └── main.go           # Web server implementation
└── README.md             # This documentation
```

## Prerequisites

- Go 1.16 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for Protocol Buffers

## Setup

1. Install required Go packages:

   ```bash
   go get google.golang.org/grpc
   go get google.golang.org/protobuf
   ```

2. Generate Go code from proto files:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       api/message.proto
   ```

## Running the Application

1. Start the gRPC server (from the root directory):

   ```bash
   go run server/main.go
   ```

2. Start the web server (from the root directory):

   ```bash
   go run web/main.go
   ```

3. Access the web interface at `http://localhost:8081`

## Features

### Submit Message

- Submit a message with a name and content
- Messages are stored on the gRPC server
- Real-time request/response display

### Get Message

- Retrieve messages by name
- View message content and metadata
- Error handling for non-existent messages

## API Endpoints

### Web Server (Port 8081)

- `GET /`: Web interface
- `POST /submit`: Submit a new message
  ```json
  {
    "name": "string",
    "message": "string"
  }
  ```
- `GET /get/{name}`: Retrieve a message by name

### gRPC Server (Port 8080)

- `SubmitMessage`: Submit a new message
- `GetMessage`: Retrieve a message by name

## Testing with Postman

Postman supports gRPC testing through its gRPC client. Here's how to set it up:

1. **Install Postman**

   - Download and install Postman from [postman.com](https://www.postman.com/downloads/)
   - Ensure you have the latest version that supports gRPC

2. **Create a New gRPC Request**

   - Click "New" → "gRPC Request"
   - Enter the server URL: `localhost:8080`
   - **Important**: Click on the "Server Reflection" toggle to turn it OFF
   - Click "Next"

3. **Import Proto File**

   - Click "Import" in the gRPC request window
   - Select "Import from File"
   - Choose your `api/message.proto` file
   - Click "Import"

4. **Test SubmitMessage**

   - Select the `SubmitMessage` method
   - In the request body, enter:
     ```json
     {
       "name": "test",
       "message": "Hello from Postman"
     }
     ```
   - Click "Invoke" to send the request

5. **Test GetMessage**
   - Select the `GetMessage` method
   - In the request body, enter:
     ```json
     {
       "name": "test"
     }
     ```
   - Click "Invoke" to retrieve the message

### Troubleshooting Postman gRPC Tests

1. **Connection Issues**

   - Ensure the gRPC server is running on port 8080
   - Check if the server URL is correct
   - Verify no firewall is blocking the connection
   - Make sure TLS/SSL is turned OFF in Postman settings
   - Ensure Server Reflection is turned OFF

2. **Common Error Messages and Solutions**

   a. **"Service unavailable" or "WRONG_VERSION_NUMBER"**

   - Solution: Turn OFF TLS/SSL in Postman settings
   - Solution: Turn OFF Server Reflection
   - Solution: Verify the server is running with insecure credentials

   b. **"Failed to connect"**

   - Solution: Check if the server is running
   - Solution: Verify the port number (8080)
   - Solution: Try using `127.0.0.1:8080` instead of `localhost:8080`

   c. **"Method not found"**

   - Solution: Verify the proto file is correctly imported
   - Solution: Check if the method name matches exactly
   - Solution: Ensure the server implements the method

3. **Proto File Issues**

   - Make sure the proto file is properly formatted
   - Verify all required fields are included in the request
   - Check if the proto file matches the server implementation

4. **Request Format**
   - Ensure the request body matches the proto definition
   - Check for any missing required fields
   - Verify the data types match the proto specification

## Error Handling

The application includes comprehensive error handling for:

- Connection failures
- Invalid requests
- Missing messages
- Server errors

## Development

### Adding New Features

1. Update the proto file in `api/message.proto`
2. Regenerate the Go code
3. Update the web interface in `templates/grpc_test.html`
4. Implement the new functionality in `web/main.go`

### Testing

1. Ensure the gRPC server is running
2. Start the web server
3. Test the interface through the web browser
4. Monitor the console for any errors

## Troubleshooting

Common issues and solutions:

1. **Connection Refused**

   - Ensure the gRPC server is running on port 8080
   - Check firewall settings

2. **Template Not Found**

   - Verify the template path is correct
   - Check file permissions

3. **gRPC Errors**
   - Verify proto file compilation
   - Check server implementation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
