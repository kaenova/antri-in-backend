package utils

import (
	"antri-in-backend/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

var (
	configs   config.Config        = config.GetConfig()
	JWTconfig middleware.JWTConfig = middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		Claims:      &JWTCustomClaims{},
		SigningKey:  []byte(configs.Secret),
	}
)

func GenerateToken(nama, email, id string) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaims{
		nama,
		email,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
