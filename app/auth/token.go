package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type InputToken struct {
	PublicID uuid.UUID
	Name     string
	Email    string
}

func GenerateToken(ctx context.Context, input InputToken) (string, error) {
	claims := jwt.MapClaims{
		"public_id": input.PublicID,
		"name":      input.Name,
		"email":     input.Email,
		"exp":       time.Now().Add(time.Minute * 60).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
