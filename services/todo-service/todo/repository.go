package todo

import (
	"fmt"
	"microGoLearn/common/model"

	"gorm.io/gorm"
)

type Repository interface {
	List(userID string) ([]*model.Todo, error)
	Add(todo *model.Todo) (*model.Todo, error)
	GetById(id string) (*model.Todo, error)
	UpdateById(id string, dataUpdate map[string]interface{}) (*model.Todo, error)
	DeleteById(id string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) List(userID string) ([]*model.Todo, error) {
	var todos []*model.Todo

	if err := r.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *repository) Add(todo *model.Todo) (*model.Todo, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *repository) GetById(id string) (*model.Todo, error) {
	var todo *model.Todo

	if err := r.db.Where("id = ?", id).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *repository) UpdateById(id string, dataUpdate map[string]interface{}) (*model.Todo, error) {
	var todo *model.Todo

	if err := r.db.Model(&todo).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("id = ?", id).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
func (r *repository) DeleteById(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&model.Todo{}).Error; err != nil {
		return "error", err
	}

	status := fmt.Sprintf("todo id %v success deleted", id)

	return status, nil
}
