package todo

import (
	"context"
	"log"
	"microGoLearn/common/model"
	"strconv"
)

type TodosServer struct {
	repository Repository
}

func NewServer(repository Repository) *TodosServer {
	return &TodosServer{repository: repository}
}

func (s *TodosServer) ListTodo(ctx context.Context, param *model.TodoUserId) (*model.TodoList, error) {
	log.Println("get service list todo")

	todos, err := s.repository.List(param.UserId)
	if err != nil {
		return nil, err
	}

	var todoList = &model.TodoList{}

	todoList.List = todos

	return todoList, nil
}

func (s *TodosServer) AddTodo(ctx context.Context, param *model.TodoCreate) (*model.Todo, error) {
	log.Println("get service Add Todo")

	userID, _ := strconv.Atoi(param.UserId)

	var newTodo = &model.Todo{
		Title:       param.Todo.Title,
		Description: param.Todo.Description,
		Category:    param.Todo.Category,
		UserId:      int32(userID),
	}

	todo, err := s.repository.Add(newTodo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodosServer) GetTodoById(ctx context.Context, param *model.TodoId) (*model.Todo, error) {
	log.Println("get service todo by ID")

	todo, err := s.repository.GetById(param.Id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodosServer) UpdateTodoById(ctx context.Context, param *model.TodoUpdate) (*model.Todo, error) {
	log.Println("get service update todo by ID")

	var dataUpdate = map[string]interface{}{}

	if param.Todo.Title != "" || len(param.Todo.Title) > 0 {
		dataUpdate["title"] = param.Todo.Title
	}

	if param.Todo.Description != "" || len(param.Todo.Description) > 0 {
		dataUpdate["description"] = param.Todo.Description
	}

	if param.Todo.Category != "" || len(param.Todo.Category) > 0 {
		dataUpdate["category"] = param.Todo.Category
	}

	todo, err := s.repository.UpdateById(param.Id, dataUpdate)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodosServer) DeleteTodoById(ctx context.Context, param *model.TodoId) (*model.Deleted, error) {
	log.Println("get service delete todo By ID")

	msg, err := s.repository.DeleteById(param.Id)
	if err != nil {
		return nil, err
	}

	message := &model.Deleted{Message: msg}

	return message, nil

}
