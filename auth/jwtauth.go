package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAuthToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})
	signedToken, err := token.SignedString([]byte("projectinidilindungi"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
