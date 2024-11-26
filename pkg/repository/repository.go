package repository

import (
	gosonglibrary "github.com/SokoloSHA/GoSongLibrary"
	"github.com/jmoiron/sqlx"
)

type Aunthorization interface {
	CreateUser(user gosonglibrary.User) (int, error)
	GetUser(username, password string) (gosonglibrary.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Aunthorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Aunthorization: NewAuthPostgres(db),
	}
}
