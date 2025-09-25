# URL Checker 🔍

A lightweight CLI tool written in Go for checking website availability and HTTP status codes. Performs concurrent URL checks with colorful output.

## Features ✨

* **Concurrent checking** - Multiple URLs checked simultaneously
* **Fast HEAD requests** - Uses HTTP HEAD method for efficiency  
* **Timeout handling** - Configurable 10-second timeout per request
* **Colorful status output** - Visual indicators for different HTTP statuses
* **Simple file-based input** - Check multiple URLs from a text file

## Installation 📦

### Prerequisites
* Go 1.16 or higher
```bash
### From Source

# Clone the repository
git clone https://github.com/gnutov95/go-url-checker.git
cd go-url-checker

# Build the binary
go build -o url-checker main.go

# Or run directly
go run main.go urls.txt
##Usage 🚀
Basic Usage
bash
# Check URLs from a file
go run main.go urls.txt

# Using the compiled binary
./url-checker urls.txt
Preparing URL File
Create a urls.txt file with URLs (one per line):

txt
https://google.com
https://github.com
https://httpbin.org/status/200
https://httpbin.org/status/404
https://httpbin.org/status/500
##Example Output 📊
bash
❌ https://invalid-domain.com - Error: dial tcp: lookup invalid-domain.com: no such host
✅ https://google.com - Status: 200
⚠️  https://httpbin.org/status/404 - Status: 404
✅ https://httpbin.org/status/200 - Status: 200
❌ https://httpbin.org/delay/10 - Error: context deadline exceeded
```
## Project Structure 🏗️
text
go-url-checker/
├── main.go          # Main application code
├── urls.txt         # Example URL list for testing
├── go.mod           # Go module file
└── README.md        # This file
## Core Functions 🔧
main()
Handles command line arguments

Coordinates URL checking process

Manages goroutines and channels

Displays results

checkUrl(url string) (int, error)
Performs HTTP HEAD requests with timeout

Returns status code and error

10-second timeout configuration

readURLs(filename string) ([]string, error)
Reads URL list from text file

Skips empty lines

Returns URL array for processing

## Building and Running 🧪
bash
# Build the application
go build -o url-checker main.go

# Run with test URLs
./url-checker urls.txt

# Or run directly with Go
go run main.go urls.txt
## Roadmap 🗺️
Potential enhancements for future versions:

Command-line flags for timeout configuration

Customizable concurrent workers count

JSON/CSV output format support

Retry mechanism for failed requests

Response time metrics

Configurable HTTP methods (GET/POST)

Bulk export results to file

Docker container support

## Contributing 🤝
Contributions are welcome! Here's how you can help:

Fork the repository

Create a feature branch (git checkout -b feature/amazing-feature)

Commit your changes (git commit -m 'Add amazing feature')

Push to the branch (git push origin feature/amazing-feature)

Open a Pull Request

## License 📄
This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer ⚠️
Use this tool responsibly and only for legitimate purposes such as:

Monitoring your own websites and services

Testing with explicit permission from resource owners

Educational and learning purposes

Please respect robots.txt files and rate limiting policies of target websites.

⭐ If you find this project useful, please give it a star!
