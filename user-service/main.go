package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
   "github.com/nastaranmotamed/go-microservices-notes/user-service/pb"

    "user-service/handler"
)

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &handler.UserHandler{})

    log.Println("UserService gRPC server is listening on port 50052...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
