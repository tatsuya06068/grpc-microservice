package grpc

import (
	"context"

	"github.com/myusername/myservice/internal/usecase"
	"github.com/myusername/myservice/proto"
)

type UserHandler struct {
	usecase usecase.IUserUsecase
	proto.UnimplementedUserServiceServer
}

func NewUserHandler(usecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := h.usecase.GetUser(req.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		Id:   user.ID,
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
