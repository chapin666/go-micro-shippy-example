package main

import (
	"context"
	"os"
	"log"
	"errors"
	"encoding/json"
	"io/ioutil"
	"github.com/micro/go-micro"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	DEFAULT_INFO_FILE = "consignment.json"
)

func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}


func main() {

	sercice := micro.NewService(micro.Name("go.micro.srv.consignment"))
	sercice.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", sercice.Client())

	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 创建托运货物
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	log.Printf("created: %t", resp.Created)

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}
	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}

}