package service

import (
	"context"
	"database/sql"

	"github.com/rizqitsani/learn-go-api/dto"
	"github.com/rizqitsani/learn-go-api/exception"
	"github.com/rizqitsani/learn-go-api/helper"
	"github.com/rizqitsani/learn-go-api/model"
	"github.com/rizqitsani/learn-go-api/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	db             *sql.DB
}

func NewUserService(repository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		userRepository: repository,
		db:             db,
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

	user = service.userRepository.Create(ctx, tx, user)

	return user
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []model.User {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	return service.userRepository.FindAll(ctx, tx)
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int) model.User {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.userRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *UserServiceImpl) Update(ctx context.Context, request dto.UpdateUserDto) model.User {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.userRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name

	user = service.userRepository.Update(ctx, tx, user)

	return user
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	_, err = service.userRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.userRepository.Delete(ctx, tx, id)
}
