package main

import (
	"log"

	item "github.com/ahang7/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/item/itemservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	itemServiceImpl := new(ItemServiceImpl)
	stockCli, err := NewStockClient("127.0.0.1:8890")
	if err != nil {
		log.Fatal(err)
	}
	itemServiceImpl.stockClient = stockCli

	svr := item.NewServer(itemServiceImpl,
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "example.shop.item",
			},
		),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
