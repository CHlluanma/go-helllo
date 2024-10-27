package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ahang7/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/item"
	"github.com/ahang7/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/item/itemservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var cli itemservice.Client

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	c, err := itemservice.NewClient("example.shop.item", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	cli = c

	hz := server.New(server.WithHostPorts("localhost:8889"))

	hz.GET("/api/item", Handler)
	if err := hz.Run(); err != nil {
		log.Fatal()
	}
}

func Handler(ctx context.Context, c *app.RequestContext) {
	req := item.NewGetItemRequest()
	req.Id = 1024
	resp, err := cli.GetIem(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, resp.String())
}
