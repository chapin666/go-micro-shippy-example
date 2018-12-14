package main

import (
	"context"
	"log"
	userPb "shippy/user-service/proto/user"

	micro "github.com/micro/go-micro"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *userPb.User) error {
	log.Println("[Picked up a new message]")
	log.Println("[Sending email to]: %v\n", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}
