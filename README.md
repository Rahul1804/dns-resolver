# DNS Resolver in Go

This project implements a simple DNS resolver in Go using the `net` package. The resolver takes a hostname as input and returns the IP addresses associated with that hostname.

## Features

- Resolves A records for a given hostname.
- Command-line interface for easy usage.
- Basic error handling and logging.

## Prerequisites

- Go 1.18 or higher installed on your machine.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/rahul1804/dns-resolver.git
    cd dns-resolver
    ```

2. Initialize the Go module:

    ```bash
    go mod init dns-resolver
    ```

## Usage

### Building the Application

To build the DNS resolver application, run the following command:

```bash
go build -o dns-resolver
./dns-resolver google.com
```
