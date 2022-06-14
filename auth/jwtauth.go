package auth

import (
	"github.com/golang-jwt/jwt"
)

func GenerateAuthToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	signedToken, err := token.SignedString([]byte("projectinidilindungi"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
