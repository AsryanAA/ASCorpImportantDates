package jwt

import (
	"ASCorpImportantDates/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const validityJWT = 42 // в секундах

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(validityJWT * time.Second)

	claims := &Claims{
		Username: user.Login,
		Password: user.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
