package user

import (
	"context"
	"errors"
	"log"
	"microGoLearn/common/model"
	"microGoLearn/services/user-service/auth"
	"strconv"
)

type UsersServer struct {
	auth       auth.Service
	repository Repository
}

func NewServer(auth auth.Service, repo Repository) *UsersServer {
	return &UsersServer{auth, repo}
}

func (s *UsersServer) Register(ctx context.Context, param *model.UserInput) (*model.UserFormat, error) {
	log.Println("get service register user")

	userCheck, err := s.repository.FindByEmail(param.Email)
	if err != nil {
		return nil, err
	}

	if userCheck.Email != "" || userCheck.Id != 0 {
		return nil, errors.New("email has been registered")
	}

	var newUser = &model.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	user, err := s.repository.Add(newUser)
	if err != nil {
		return nil, err
	}

	userID := strconv.Itoa(int(user.Id))

	token, err := s.auth.GenerateToken(userID, user.Email)

	if err != nil {
		return nil, err
	}

	userFormat := Formatter(user, token)

	return userFormat, nil

}

func (s *UsersServer) Login(ctx context.Context, param *model.UserLogin) (*model.UserFormat, error) {
	log.Println("get service login user")

	user, err := s.repository.FindByEmail(param.Email)
	if err != nil {
		return nil, err
	}

	if user.Email == "" || user.Id == 0 {
		return nil, errors.New("invalid email / password")
	}

	if user.Password != param.Password {
		return nil, errors.New("invalid email / password")
	}

	userID := strconv.Itoa(int(user.Id))

	token, err := s.auth.GenerateToken(userID, user.Email)

	if err != nil {
		return nil, err
	}

	userFormat := Formatter(user, token)

	return userFormat, nil
}
