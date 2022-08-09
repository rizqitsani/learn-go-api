package service

import (
	"context"

	"github.com/rizqitsani/learn-go-api/dto"
	"github.com/rizqitsani/learn-go-api/model"
)

type UserService interface {
	Create(ctx context.Context, request dto.CreateUserDto) model.User
	FindAll(ctx context.Context) []model.User
	FindById(ctx context.Context, id int) model.User
}
