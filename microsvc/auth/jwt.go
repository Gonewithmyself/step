package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type payload struct {
	Usr string `json:"user"`
	Psw string `json:"pswd"`
	jwt.StandardClaims
}

var secret string

func genToken(payload *payload) (string, error) {
	payload.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secret))
}

func parseToken(token string) (*payload, error) {
	pl := &payload{}
	tk, er := jwt.ParseWithClaims(token, pl, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if er != nil {
		return nil, er
	}

	return tk.Claims.(*payload), nil
}
