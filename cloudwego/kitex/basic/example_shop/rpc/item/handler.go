package main

import (
	"context"
	"log"

	item "github.com/CHlluanma/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/item"
	"github.com/CHlluanma/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/stock"
	"github.com/CHlluanma/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/stock/stockservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct {
	stockClient stockservice.Client
}

func NewStockClient(addr string) (stockservice.Client, error) {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	return stockservice.NewClient("example.shop.stock", client.WithResolver(r))
}

// GetIem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetIem(ctx context.Context, request *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	resp = item.NewGetItemResponse()
	resp.Item = item.NewItem()
	resp.Item.Id = request.GetId()
	resp.Item.Title = "kitex"
	resp.Item.Description = "kitex is an excellent framework!"

	stockReq := stock.NewGetItemStockRequest()
	stockReq.ItemId = request.GetId()
	stockResp, err := s.stockClient.GetItemStock(context.Background(), stockReq)
	if err != nil {
		log.Println(err)
		stockResp.Stock = 0
	}
	resp.Item.Stock = stockResp.Stock
	return
}
