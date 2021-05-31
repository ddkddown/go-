package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"ddk/myproto"

	"net/http"

	"google.golang.org/grpc"
)

type RegisReq struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Port    int      `json:"port"`
	Tags    []string `json:"tags"`
}

func main() {
	grpcServer := grpc.NewServer()
	myproto.RegisterHelloServiceServer(grpcServer, new(HelloServiceImp))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	regis := &RegisReq{
		Id:      "grpc_test",
		Name:    "grpc",
		Address: "192.168.10.210",
		Port:    1234,
		Tags:    []string{"grpc_server"},
	}

	client := &http.Client{}

	json, err := json.Marshal(regis)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, "http://192.168.10.210:8500/v1/agent/service/register", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	grpcServer.Serve(lis)
}
