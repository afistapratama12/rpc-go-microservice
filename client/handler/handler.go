package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"microGoLearn/common/config"
	"microGoLearn/common/model"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var (
	user     = ServiceUser()
	todo     = ServiceTodo()
	conn     = config.RedisConn()
	redisKey = "todo:"
)

type clientHandler struct{}

func NewClientHandler() *clientHandler {
	return &clientHandler{}
}

func (h *clientHandler) Login(c *gin.Context) {
	var input *model.UserLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := user.Login(context.Background(), input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *clientHandler) Register(c *gin.Context) {
	var input *model.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := user.Register(context.Background(), input)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func (h *clientHandler) ListTodo(c *gin.Context) {
	userID := c.MustGet("currentUserId").(string)

	fmt.Println(userID)

	if userID == "" {
		c.JSON(401, gin.H{"error": "unauthorize user"})
		return
	}

	reply, err := redis.Bytes(conn.Do("GET", redisKey+userID))
	if err == nil {
		log.Println("get data from redis")
		todos := &model.TodoList{}
		json.Unmarshal(reply, &todos)

		c.JSON(200, todos)
		return
	}

	getTodos := &model.TodoUserId{UserId: userID}

	todos, err := todo.ListTodo(context.Background(), getTodos)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	dataByte, _ := json.Marshal(todos)

	conn.Do("SET", redisKey+userID, dataByte)

	c.JSON(200, todos)
}

func (h *clientHandler) AddTodo(c *gin.Context) {
	userID := c.MustGet("currentUserId").(string)

	if userID == "" {
		c.JSON(401, gin.H{"error": "unauthorize user"})
		return
	}

	var input *model.TodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var addTodo = &model.TodoCreate{UserId: userID}

	addTodo.Todo = input

	todo, err := todo.AddTodo(context.Background(), addTodo)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	conn.Do("DEL", redisKey+userID)
	log.Println("deleting data from redis")

	c.JSON(201, todo)
}

func (h *clientHandler) GetTodoById(c *gin.Context) {
	userID := c.MustGet("currentUserId").(string)

	if userID == "" {
		c.JSON(401, gin.H{"error": "unauthorize user"})
		return
	}

	todoId := c.Params.ByName("todo_id")

	getTodoId := &model.TodoId{Id: todoId}

	todo, err := todo.GetTodoById(context.Background(), getTodoId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todo)
}

func (h *clientHandler) UpdateTodoById(c *gin.Context) {
	userID := c.MustGet("currentUserId").(string)

	if userID == "" {
		c.JSON(401, gin.H{"error": "unauthorize user"})
		return
	}

	todoId := c.Params.ByName("todo_id")

	var input *model.TodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updateTodo := &model.TodoUpdate{Id: todoId}
	updateTodo.Todo = input

	todo, err := todo.UpdateTodoById(context.Background(), updateTodo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	conn.Do("DEL", redisKey+userID)
	log.Println("deleting data from redis")

	c.JSON(200, todo)
}

func (h *clientHandler) DeleteTodoById(c *gin.Context) {
	userID := c.MustGet("currentUserId").(string)

	if userID == "" {
		c.JSON(401, gin.H{"error": "unauthorize user"})
		return
	}

	todoId := c.Params.ByName("todo_id")

	deleteTodo := &model.TodoId{Id: todoId}

	deleted, err := todo.DeleteTodoById(context.Background(), deleteTodo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	conn.Do("DEL", redisKey+userID)
	log.Println("deleting data from redis")

	c.JSON(200, deleted)
}
