package grpcx

import (
	"context"
	"estudo-test/api/grpc/example"
	"estudo-test/internal/service"
)

type UserServiceServer struct {
	example.UnimplementedUserServiceServer
	UserService service.UserService
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *example.CreateUserRequest) (*example.CreateUserResponse, error) {
	id, err := s.UserService.CreateUser(req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &example.CreateUserResponse{Id: id}, nil
}

func (s *UserServiceServer) GetUserById(ctx context.Context, req *example.GetUserByIdRequest) (*example.GetUserByIdResponse, error) {
	user, err := s.UserService.GetUsersById(req.Id)
	if err != nil {
		return nil, err
	}
	return &example.GetUserByIdResponse{Id: user.ID, Name: user.Nome, Email: user.Email}, nil
}
