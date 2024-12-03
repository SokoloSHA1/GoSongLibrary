package service

import (
	gosonglibrary "github.com/SokoloSHA/GoSongLibrary"
	"github.com/SokoloSHA/GoSongLibrary/pkg/repository"
)

type Aunthorization interface {
	GenerateToken(ip string) (string, error)
	GenerateRefreshToken(guid, email, ip string) (string, error)
	CheckRefreshToken(refresh_token string) (gosonglibrary.Users_token, error)
	SendMail() error
}

type Service struct {
	Aunthorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Aunthorization: NewAuthService(repos.Aunthorization),
	}
}
