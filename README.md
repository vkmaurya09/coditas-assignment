## Project Overview

A RESTful API service that validates Indian PAN (Permanent Account Number) card details along with other user information.

This service is part of the Coditas assignment that demonstrates best practices in Go programming including:
- RESTful API design
- Input validation
- Dependency injection
- Middleware implementation
- Logging
- Unit testing

## Getting Started

To get started with this project, clone the repository to your local machine:

```bash
git clone git@github.com:vkmaurya09/coditas-assignment.git
```

## Prerequisites

Make sure you have the following installed on your system:

- Git
- Go (Golang)

## Installation

Navigate to the project directory and install the dependencies:

```bash
cd coditas-assignment
go mod tidy
```

## Usage

Run the project using the following command:

```bash
go run cmd/pancard-service/main.go
```

You can test the API using the following `curl` command:

```bash
curl --location 'localhost:8080/submit' \
--header 'Content-Type: application/json' \
--data '{
    "name": "vinay",
    "email": "vkm@gmail.com",
    "pan": "AAAAA1234A",
    "mobile": "2345432123"
}'
```