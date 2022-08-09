package repository

import (
	"context"
	"database/sql"

	"github.com/rizqitsani/learn-go-api/model"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	query := "INSERT INTO users (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, user.Name)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int(id)
	return user
}
