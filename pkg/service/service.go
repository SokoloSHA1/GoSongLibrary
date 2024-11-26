package service

import (
	gosonglibrary "github.com/SokoloSHA/GoSongLibrary"
	"github.com/SokoloSHA/GoSongLibrary/pkg/repository"
)

type Aunthorization interface {
	CreateUser(user gosonglibrary.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Aunthorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Aunthorization: NewAuthService(repos.Aunthorization),
	}
}
