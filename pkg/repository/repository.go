package repository

import (
	gosonglibrary "github.com/SokoloSHA/GoSongLibrary"
	"github.com/jmoiron/sqlx"
)

type Aunthorization interface {
	CreateUserToken(guid, email, ip, refresh_token string) error
	GetUserToken(guid string) (gosonglibrary.Users_token, error)
	GetRefreshToken(refresh_token string) (gosonglibrary.Users_token, error)
	DeleteRefreshToken(user_token gosonglibrary.Users_token) error
}

type Repository struct {
	Aunthorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Aunthorization: NewAuthPostgres(db),
	}
}
