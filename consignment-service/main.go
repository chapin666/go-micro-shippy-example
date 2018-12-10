package main

import (
	"os"
	"log"
	"github.com/micro/go-micro"
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
)

const (
	DEFAULT_HOST = "localhost:27017"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_HOST
	}

	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}
	defer session.Close()

	server := micro.NewService(
		 micro.Name("go.micro.srv.consignment"),
		 micro.Version("latest"),
	 )

	 server.Init()

	 vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	
	 pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vClient})

	 if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	 }

}