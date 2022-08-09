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

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	query := "SELECT * FROM users"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) model.User {
	var user model.User

	query := "SELECT * FROM users WHERE id = ?"
	err := tx.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Name)
	if err != nil {
		panic(err)
	}

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	query := "UPDATE users SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Name, user.Id)
	if err != nil {
		panic(err)
	}

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}
