# Network Speed Test

A simple web application to test network download and upload speeds, as well as ping.

## Features

- Measure download speed
- Measure upload speed
- Display ping
- Show test time and server IP
- Dark mode toggle

## Prerequisites

- Go (version X.X or higher)
- [Any other dependencies]

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/network-speed-test.git
   cd network-speed-test
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```

2. Open a web browser and navigate to `http://localhost:8080` (or whatever port you've configured).

3. Click the "Run Speed Test" button to start a new test.

4. Toggle between light and dark modes using the button in the bottom right corner.

## Configuration

[Add any configuration details here, such as how to change the port or customize the speed test]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements
This project makes use of the following libraries:

- [go-speedtest](https://github.com/showwin/speedtest-go) - Go client library for speedtest.net
- [gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [html/template](https://pkg.go.dev/html/template) - Go's built-in HTML templating package
- [net/http](https://pkg.go.dev/net/http) - Go's HTTP client and server implementations
- [encoding/json](https://pkg.go.dev/encoding/json) - Go's JSON encoding and decoding package
- [time](https://pkg.go.dev/time) - Go's time package for handling dates and durations