# Golang GRPC User Service With Search

This project encompasses a Golang gRPC service designed with specialized functionalities to effectively manage user details, alongside a robust search feature.


# Overview

The service acts as a simulated database, maintaining a list of user details within a variable. It provides gRPC endpoints for fetching user details based on a user ID, retrieving a list of user details based on a list of user IDs, and implementing a search functionality to find user details based on specific criteria(height,name,id, e.t.c)

## Design  Patterns  
- ### Repository Pattern:
The Repository Pattern is utilized within the db package to abstract data access operations. This pattern offers several advantages:

- `Separation of Concerns:` It separates data access logic from business logic, promoting a cleaner and more maintainable codebase.
Clean API: Provides a clear and consistent API for interacting with data, abstracting the underlying data storage details.

- `Flexibility:` Enables easy switching or extension of data sources without impacting the application's core logic.

- `Scalability:` Ensures a clear separation of concerns between data access and other application components, facilitating scalability and easier maintenance.
- ###  Server Pattern:
The Server Pattern is employed within the main package to initialize and start the gRPC server. Key points about this pattern include:

- `Initialization:` Involves creating a server instance, registering services to handle requests, and configuring it to listen for incoming connections

- `Scalability and Maintainability:` Facilitates the development of scalable and maintainable server applications by structuring the handling of incoming requests.

- `Structured Request Handling:` Allows for the handling of incoming requests in a structured and organized manner, improving code readability and maintainability.
## Features

- Search for user details using various criteria such as city, phone number, or marital status
- Access user details using their unique user ID
- Obtain a collection of user details by providing a list of user IDs
- Containerized with Docker for seamless deployment
- Unit testing coverage for endpoints
## Prerequisites

- Go 1.21.4
- Protocol Buffers compiler (protoc)
- Docker (for containerization)

## Installation

### Clone the Repository

```sh
git clone https://github.com/akshaykajithkumar/Totality-Corp-Assignment

cd Totality-Corp-Assignment

```


## Usage

### Run the Application
 
```sh
go run cmd/server/main.go 
```
## Accessing gRPC Endpoints

You can use any gRPC client to interact with the service. Below are the examples using `postman` to interact with the various endpoints.

### Fetch User Details by User ID

Fetch user details by ID:

```sh
{
    "id": "1"
}

```
### Retrieve a List of User Details by a List of User IDs
```sh
{
    "pageNumber": 0,
    "pageSize": 3,
    "ids": ["1", "2", "3"]
}
```

### Search User Details Based On Specific Criteria
```sh
example : search user details by name

{
  "page_number": 0,
  "page_size": 1,
  "filters": {
    "name": "Anjali Menon"
  }
}


```

## Testing
Run the unit tests using the following command:

```sh
go test ./...
```
## Dockerization
### Build the Docker Image
```sh
docker build -t golang-grpc-service .
```
### Run the Docker Container
```sh
docker run -p 50051:50051 golang-grpc-service
```

