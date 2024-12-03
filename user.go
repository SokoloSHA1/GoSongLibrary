package gosonglibrary

type Users_token struct {
	Id            int    `json:"id" db:"id"`
	GUID          string `json:"guid" db:"guid"`
	Email         string `json:"email" db:"email"`
	Ip            string `json:"ip" db:"ip"`
	Refresh_token string `json:"refresh_token" db:"refresh_token"`
}
