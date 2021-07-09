package main

import (
	"log"
	"microGoLearn/common/config"
	"microGoLearn/common/model"
	"microGoLearn/services/todo-service/todo"
	"net"

	"google.golang.org/grpc"
)

var (
	db       = config.DBConn(config.DB_TODO_NAME)
	todoRepo = todo.NewRepository(db)
	todoSrv  = todo.NewServer(todoRepo)
)

func main() {
	db.AutoMigrate(&model.Todo{})

	srv := grpc.NewServer()

	model.RegisterTodosServer(srv, todoSrv)

	log.Println("server RPC running at", config.TODO_SERVER_PORT)

	l, err := net.Listen("tcp", config.TODO_SERVER_PORT)
	if err != nil {
		log.Fatalf("error running %s : %v", config.TODO_SERVER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
