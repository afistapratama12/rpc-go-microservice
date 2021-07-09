package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const (
	key = "SECRET_KEY"
)

type Service interface {
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
