package user

import (
	"microGoLearn/common/model"

	"gorm.io/gorm"
)

type Repository interface {
	Add(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Add(user *model.User) (*model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (*model.User, error) {
	var user *model.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
