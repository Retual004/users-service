// internal/transport/grpc/handler.go
package grpc

import (
	"context"

	userpb "github.com/Retual004/project-protos/proto/user"
	"github.com/Retual004/users-service/internal/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Handler реализует сгенерированный интерфейс userpb.UserServiceServer.
type Handler struct {
	svc *user.UserService
	userpb.UnimplementedUserServiceServer
}

// NewHandler конструктор.
func NewHandler(svc *user.UserService) *Handler {
	return &Handler{svc: svc}
}

// CreateUser
func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// Маппим из protobuf в внутреннюю модель
	u, err := h.svc.CreateUser(user.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	// Ответ
	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       uint32(u.ID),
			Email:    u.Email,
			Password: u.Password,
		},
	}, nil
}

// GetUser
func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.GetUserResponse, error) {
	u, err := h.svc.GetUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{
	User: &userpb.User{
		Id : uint32(u.ID),
		Email: u.Email,
		Password: u.Password,
	},
	}, nil
}

// UpdateUser
func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	orig := req.User
	u, err := h.svc.UpdateUserByID(uint(orig.Id), user.User{
		Email:    orig.Email,
		Password: orig.Password,
	})
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:       uint32(u.ID),
			Email:    u.Email,
			Password: u.Password,
		},
	}, nil
}

// DeleteUser
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := h.svc.DeleteUserByID(uint(req.Id)); err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{}, nil
}


func (h *Handler) ListUsers(ctx context.Context, _ *emptypb.Empty) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.ListUsers()
	if err != nil {
	  return nil, err
	}
	var pbUsers []*userpb.User
	for _, u := range users {
	  pbUsers = append(pbUsers, &userpb.User{
		Id:       uint32(u.ID),
		Email:    u.Email,
		Password: u.Password,
	  })
	}
	return &userpb.ListUsersResponse{Users: pbUsers}, nil
  }