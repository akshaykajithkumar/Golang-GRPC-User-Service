package main

import (
	"flag"
	"fmt"
	"log"
	"main/db"
	"main/service"
	"net"

	api "main/proto/api"

	"google.golang.org/grpc"
)

var (
	serverPort = flag.Int("port", 50051, "The server port")
)

// main function initializes and starts the gRPC server
func main() {
	flag.Parse()

	// Listens on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *serverPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Creates a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Creates a new instance of UserService with a UserRepository
	userService := service.NewUserService(db.NewUserRepository())

	// Registering the UserService with the gRPC server
	api.RegisterUserServiceServer(grpcServer, userService)

	// Log server address
	log.Printf("Server listening at %v", listener.Addr())
	//
	// Serving incoming gRPC requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
