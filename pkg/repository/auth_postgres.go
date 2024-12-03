package repository

import (
	"fmt"

	gosonglibrary "github.com/SokoloSHA/GoSongLibrary"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUserToken(guid string) (gosonglibrary.Users_token, error) {
	var user_token gosonglibrary.Users_token

	check := fmt.Sprintf("SELECT * FROM %s WHERE guid=$1", usersToken)

	err := r.db.Get(&user_token, check, guid)

	return user_token, err
}

func (r *AuthPostgres) CreateUserToken(guid, email, ip, refresh_token string) error {
	query := fmt.Sprintf("INSERT INTO %s (guid, email, ip, refresh_token) values ($1, $2, $3, $4) RETURNING id", usersToken)
	_, err := r.db.Query(query, guid, email, ip, refresh_token)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthPostgres) GetRefreshToken(refresh_token string) (gosonglibrary.Users_token, error) {
	var user_token gosonglibrary.Users_token

	query := fmt.Sprintf("SELECT * FROM %s WHERE refresh_token=$1", usersToken)

	err := r.db.Get(&user_token, query, refresh_token)

	return user_token, err
}

func (r *AuthPostgres) DeleteRefreshToken(user_token gosonglibrary.Users_token) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersToken)

	_, err := r.db.Exec(query, user_token.Id)

	return err
}
