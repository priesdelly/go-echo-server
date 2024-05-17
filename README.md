# REQ-RES - Go HTTP Request Echo Server
This is a simple Go application that demonstrates how to create an HTTP server that echoes back the request details (method, URL, headers, and body) in JSON format.

**Features:**

- Handles all HTTP methods (GET, POST, PUT, etc.)
- Uses only native Go libraries for dependencies

**Running the Application:**

1. Save the code as `main.go`.
2. Run the application using: `go run main.go`
3. The server will listen on port 8080.

**Testing:**

You can test the application using tools like curl or Postman. Send a request to `http://localhost:8080` (or the appropriate URL if running on a different machine) with any HTTP method and optional body content. The response will be a JSON object containing the request details.


## Build Requirements:

Go version 1.22 or later (instructions on installing Go can be found at https://go.dev/doc/install)

**Installation:**

Clone this repository:
``` bash
git clone https://github.com/priesdelly/req-res.git
```
Navigate to the project directory:
``` bash
cd req-res
```
Run go mod download to download dependencies:
```bash
go mod download
```
Build the application:
```bash
go build -o main
```

## Running the Application:

Execute the built binary:
```bash
./main
```
**Example Usage:**
```bash
curl --location 'http://localhost:8080/call/req-res'
```

**License:**

This application is distributed under the terms of the [BSD-3-Clause License](https://opensource.org/licenses/BSD-3-Clause).  For more information, see the LICENSE file.

Author(s):
- [Priesdelly](https://github.com/priesdelly)

