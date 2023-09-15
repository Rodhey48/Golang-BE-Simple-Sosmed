package middlewares

import (
	"os"
	"simple_sosmed/models/users/entity"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(user entity.User) string {
	claims := &jwtCustomClaims{
		user.Id,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(os.Getenv("PRIVATE_KEY_JWT")))

	return t
}
