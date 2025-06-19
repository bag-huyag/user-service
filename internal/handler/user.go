package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/bag-huyag/user-service/internal/kafka"
	pb "github.com/bag-huyag/user-service/proto/gen"
	"github.com/google/uuid"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	producer *kafka.Producer
}

func NewUserHandler(p *kafka.Producer) *UserHandler {
	return &UserHandler{producer: p}
}

func (h *UserHandler) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	id := uuid.New().String()
	user := &pb.User{
		Id:    id,
		Name:  in.Name,
		Email: in.Email,
	}

	h.producer.SendUserEvent(kafka.UserEvent{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Type:  "created",
	})

	return user, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	if in.Id == "" {
		return nil, errors.New("missing user ID")
	}

	h.producer.SendUserEvent(kafka.UserEvent{
		ID:    in.Id,
		Name:  in.Name,
		Email: in.Email,
		Type:  "updated",
	})

	return in, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, in *pb.UserId) (*pb.Empty, error) {
	if in.Id == "" {
		return nil, errors.New("missing user ID")
	}

	h.producer.SendUserEvent(kafka.UserEvent{
		ID:   in.Id,
		Type: "deleted",
	})

	return &pb.Empty{}, nil
}

// GetUser — временная заглушка
func (h *UserHandler) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	return nil, fmt.Errorf("GetUser not implemented in user-service")
}

// GetUsers — временная заглушка
func (h *UserHandler) GetUsers(ctx context.Context, _ *pb.Empty) (*pb.UserList, error) {
	return nil, fmt.Errorf("GetUsers not implemented in user-service")
}
