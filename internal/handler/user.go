package handler

import (
	"context"

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

//TODO Другие методы (Update, Delete...)
