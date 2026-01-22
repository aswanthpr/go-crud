package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenSecret  = []byte("ACCESS_SECRET_KEY")
	refreshTokenSecret = []byte("REFRESH_SECRET_KEY")
)

type TokenClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint, email string) (string, error) {
	Claims := TokenClaim {
		UserID:  userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	return token.SignedString(accessTokenSecret)
}