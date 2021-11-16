package utils

import (
	"antri-in-backend/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaimsAdmin struct {
	Nama  string `json:"nama"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

type JWTCustomClaimsPengantri struct {
	Id        string `json:"id"`
	Nama      string `json:"nama"`
	NoAntrian int    `json:"no_antrian"`
	IdAntrian string `json:"id_antrian"`
	jwt.StandardClaims
}

var (
	configs        config.Config        = config.GetConfig()
	JWTconfigAdmin middleware.JWTConfig = middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		Claims:      &JWTCustomClaimsAdmin{},
		SigningKey:  []byte(configs.Secret),
	}
	JWTconfigPengantri middleware.JWTConfig = middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		Claims:      &JWTCustomClaimsPengantri{},
		SigningKey:  []byte(configs.Secret),
	}
)

func GenerateTokenAdmin(nama, role, email, id string) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaimsAdmin{
		nama,
		role,
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

func GenerateTokenPengantri(id, nama, idAntrian string, noAntrian int) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaimsPengantri{
		id,
		nama,
		noAntrian,
		idAntrian,
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

func JWTTokenFromStringPengantri(key string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(key, &JWTCustomClaimsPengantri{}, func(t *jwt.Token) (interface{}, error) {
		conf := config.GetConfig()
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(conf.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
