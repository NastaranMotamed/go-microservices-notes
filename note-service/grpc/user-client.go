package grpc

import (
    "context"
    "time"

    "google.golang.org/grpc"
    "note-service/pb"
)

type UserClient struct {
    client pb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
    return &UserClient{client: pb.NewUserServiceClient(conn)}
}

func (u *UserClient) GetUser(id string) (*pb.UserResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    res, err := u.client.GetUser(ctx, &pb.UserRequest{Id: id})
    if err != nil {
        return nil, err
    }

    return res, nil
}
