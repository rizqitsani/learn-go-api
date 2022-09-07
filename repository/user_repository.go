package repository

import (
	"context"
	"database/sql"

	"github.com/rizqitsani/learn-go-api/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.User) model.User
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
	FindById(ctx context.Context, tx *sql.Tx, id int) (model.User, error)
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, id int)
}
