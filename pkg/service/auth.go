package service

import (
	"fmt"
	"math/rand"
	"time"

	gosonglibrary "github.com/SokoloSHA/GoSongLibrary"
	"github.com/SokoloSHA/GoSongLibrary/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "asdag23jfkhj13jh23hjhj"
	signingKey = "asdgw4334F$F$%$F"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Ip string `json:"ip"`
}

type AuthService struct {
	repo repository.Aunthorization
}

func NewAuthService(repo repository.Aunthorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(ip string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ip,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateRefreshToken(guid, email, ip string) (string, error) {
	user_token, err := s.repo.GetUserToken(guid)
	if err == nil {
		err := s.repo.DeleteRefreshToken(user_token)
		if err != nil {
			return "", err
		}
	}

	b := make([]byte, 32)

	f := rand.NewSource(time.Now().Unix())
	r := rand.New(f)

	_, err = r.Read(b)
	if err != nil {
		return "", err
	}

	refresh_token, err := bcrypt.GenerateFromPassword([]byte(b), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	err = s.repo.CreateUserToken(guid, email, ip, fmt.Sprintf("%x", refresh_token))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", refresh_token), nil

}

func (s *AuthService) CheckRefreshToken(refresh_token string) (gosonglibrary.Users_token, error) {
	user_token, err := s.repo.GetRefreshToken(refresh_token)

	return user_token, err
}

func (s *AuthService) SendMail() error {
	fmt.Printf("Отправка сообщения на Email")
	return nil
}
