package grpchandler

import (
	"context"

	pb "github.com/armedi/learn-go/grpc/user"
	"github.com/armedi/learn-go/user"
)

type userServer struct {
	pb.UnimplementedUserServer
	userService user.Service
}

// NewUserHandler creates an object that represent pb.UserServer Interface
func NewUserHandler(us user.Service) pb.UserServer {
	return &userServer{
		userService: us,
	}
}

func (us *userServer) Register(ctx context.Context, form *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	data := &user.RegisterRequest{
		Email:    form.GetEmail(),
		Password: form.GetPassword(),
	}
	if err := us.userService.Register(data); err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (us *userServer) Login(ctx context.Context, form *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := us.userService.Login(&user.LoginRequest{
		Email:    form.GetEmail(),
		Password: form.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}
