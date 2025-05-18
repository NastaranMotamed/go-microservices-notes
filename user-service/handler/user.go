package handler

import (
    "context"

    "github.com/nastaranmotamed/go-microservices-notes/user-service/pb"

)

type UserHandler struct {
    pb.UnimplementedUserServiceServer
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
    return &pb.UserResponse{
        Id:   req.GetId(),
        Name: "User " + req.GetId(),
    }, nil
}
