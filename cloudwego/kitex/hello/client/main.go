package main

import (
	"context"
	"log"
	"time"

	"github.com/chhz0/go-hello/cloudwego/kitex/hello/kitex_gen/api"
	"github.com/chhz0/go-hello/cloudwego/kitex/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := hello.NewClient("hello", client.WithHostPorts(":8888"))
	if err != nil {
		panic(err)
	}
	for {
		req := &api.Request{Message: "hello kitex"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)

		addReq := &api.AddRequest{First: 1, Second: 2}
		addResp, err := client.Add(context.Background(), addReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(addResp)
		time.Sleep(time.Second)
	}
}
