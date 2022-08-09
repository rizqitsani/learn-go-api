package repository

import (
	"context"
	"database/sql"

	"github.com/rizqitsani/learn-go-api/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.User) model.User
}
