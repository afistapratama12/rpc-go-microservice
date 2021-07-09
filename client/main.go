package main

import (
	"microGoLearn/client/auth"
	"microGoLearn/client/handler"

	"github.com/gin-gonic/gin"
)

var (
	clientHandler = handler.NewClientHandler()
	authService   = auth.NewService()
	middleware    = handler.Middleware(authService)
)

func main() {
	r := gin.Default()

	r.POST("/login", clientHandler.Login)
	r.POST("/register", clientHandler.Register)
	r.GET("/todos", middleware, clientHandler.ListTodo)
	r.POST("/todos", middleware, clientHandler.AddTodo)
	r.GET("/todos/:todo_id", middleware, clientHandler.GetTodoById)
	r.PUT("/todos/:todo_id", middleware, clientHandler.UpdateTodoById)
	r.DELETE("/todos/:todo_id", middleware, clientHandler.DeleteTodoById)

	r.Run()
}
