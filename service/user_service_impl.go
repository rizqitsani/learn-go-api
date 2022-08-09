package service

import (
	"context"
	"database/sql"

	"github.com/rizqitsani/learn-go-api/dto"
	"github.com/rizqitsani/learn-go-api/helper"
	"github.com/rizqitsani/learn-go-api/model"
	"github.com/rizqitsani/learn-go-api/repository"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	db         *sql.DB
}

func NewUserService(repository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		repository: repository,
		db:         db,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.CreateUserDto) model.User {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user := model.User{
		Name: request.Name,
	}

	user = service.repository.Create(ctx, tx, user)

	return user
}
