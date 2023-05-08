package middleware

import (
	"kasir/cafe/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string, password string) (string, error) {

	// payload
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return + signature
	return token.SignedString([]byte(constant.SECRET_JWT))

}
