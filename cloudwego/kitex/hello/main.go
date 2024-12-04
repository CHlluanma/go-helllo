package main

import (
	"log"

	"github.com/chhz0/go-hello/cloudwego/kitex/hello/kitex_gen/api/hello"
)

func main() {
	svr := hello.NewServer(new(HelloImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
