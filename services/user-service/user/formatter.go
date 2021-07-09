package user

import "microGoLearn/common/model"

func Formatter(user *model.User, token string) *model.UserFormat {
	return &model.UserFormat{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: token,
	}
}
