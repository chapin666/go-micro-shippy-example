package main

import (
	"fmt"
	"log"
	pb "shippy/user-service/proto/user"

	micro "github.com/micro/go-micro"
)

func main() {

	db, err := CreateConnection()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	fmt.Printf("%+v\n", db)

	defer db.Close()

	repo := &UserRepository{db}
	db.AutoMigrate(&pb.User{})

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	s.Init()

	t := TokenService{repo}
	pb.RegisterUserServiceHandler(s.Server(), &handler{repo, &t})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}
