package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(req string, reply *string) error {
	return p.Client.Call(HelloServiceName+".CallHello", req, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("rpc", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
