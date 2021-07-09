package main

import (
	"log"
	"microGoLearn/common/config"
	"microGoLearn/common/model"
	"microGoLearn/services/user-service/auth"
	"microGoLearn/services/user-service/user"

	"net"

	"google.golang.org/grpc"
)

var (
	db          = config.DBConn(config.DB_USER_NAME)
	userRepo    = user.NewRepository(db)
	authService = auth.NewService()
	userSrv     = user.NewServer(authService, userRepo)
)

func main() {
	db.AutoMigrate(&model.User{})

	srv := grpc.NewServer()

	model.RegisterUsersServer(srv, userSrv)

	log.Println("starting RPC server at", config.USER_SERVER_PORT)

	l, err := net.Listen("tcp", config.USER_SERVER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s : %v", config.USER_SERVER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
