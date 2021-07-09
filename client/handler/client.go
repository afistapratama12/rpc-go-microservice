package handler

import (
	"log"
	"microGoLearn/common/config"
	"microGoLearn/common/model"

	"google.golang.org/grpc"
)

func ServiceTodo() model.TodosClient {
	port := config.TODO_SERVER_PORT

	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewTodosClient(conn)
}

func ServiceUser() model.UsersClient {
	port := config.USER_SERVER_PORT

	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}
