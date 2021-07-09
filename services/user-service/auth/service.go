package auth

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	key = "SECRET_KEY"
)

type Service interface {
	GenerateToken(userID, email string) (string, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID, email string) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
