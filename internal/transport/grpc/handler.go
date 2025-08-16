package grpc

import (
	"context"

	userpb "github.com/Kilril312/project-protos/proto/user"
	"github.com/Kilril312/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u := &user.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.svc.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := h.svc.GetUser(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	pUser := userpb.User{
		Id:    uint32(user.ID),
		Email: user.Email,
	}

	return &userpb.GetUserResponse{
		User: &pUser,
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.User, error) {
	updUser := &user.User{
		ID:       uint(req.Id),
		Email:    req.Newemail,
		Password: req.Newpassword,
	}

	if err := h.svc.UpdateUser(ctx, updUser); err != nil {
		return nil, err
	}

	updatedUser, err := h.svc.GetUser(ctx, uint(req.Id))
	if err != nil {
		return nil, nil
	}

	return &userpb.User{
		Id:    uint32(updatedUser.ID),
		Email: updatedUser.Email,
	}, nil
}
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := h.svc.DeleteUser(ctx, uint(req.Id)); err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Success: true}, nil

}

func (h *Handler) ListUsers(ctx context.Context, _ *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.ListUser(ctx)
	if err != nil {
		return nil, err
	}

	list := &userpb.ListUsersResponse{
		Users: make([]*userpb.User, 0, len(users)),
	}

	for _, u := range users {
		list.Users = append(list.Users, &userpb.User{
			Id:       uint32(u.ID),
			Email:    u.Email,
			Password: u.Password,
		})
	}

	return list, nil
}
