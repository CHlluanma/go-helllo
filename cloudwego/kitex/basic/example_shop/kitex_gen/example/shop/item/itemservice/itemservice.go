// Code generated by Kitex v0.11.3. DO NOT EDIT.

package itemservice

import (
	"context"
	"errors"
	item "github.com/chhz0/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/item"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetIem": kitex.NewMethodInfo(
		getIemHandler,
		newItemServiceGetIemArgs,
		newItemServiceGetIemResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	itemServiceServiceInfo                = NewServiceInfo()
	itemServiceServiceInfoForClient       = NewServiceInfoForClient()
	itemServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return itemServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return itemServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return itemServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ItemService"
	handlerType := (*item.ItemService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "item",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func getIemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*item.ItemServiceGetIemArgs)
	realResult := result.(*item.ItemServiceGetIemResult)
	success, err := handler.(item.ItemService).GetIem(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newItemServiceGetIemArgs() interface{} {
	return item.NewItemServiceGetIemArgs()
}

func newItemServiceGetIemResult() interface{} {
	return item.NewItemServiceGetIemResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetIem(ctx context.Context, request *item.GetItemRequest) (r *item.GetItemResponse, err error) {
	var _args item.ItemServiceGetIemArgs
	_args.Request = request
	var _result item.ItemServiceGetIemResult
	if err = p.c.Call(ctx, "GetIem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
