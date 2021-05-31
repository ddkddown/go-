package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"ddk/myproto"

	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
)

func main() {
	url := "http://192.168.10.210:8500/v1/health/service/grpc"
	reqClient := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := reqClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	address := gjson.Get(string(body), "#.Service.Address").Array()[0].String() + ":" +
		gjson.Get(string(body), "#.Service.Port").Array()[0].String()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := myproto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &myproto.Test{Name: "ddk"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply.GetName())
}
