// Code generated by Kitex v0.11.3. DO NOT EDIT.

package itemservice

import (
	"context"
	item "github.com/CHlluanma/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen/example/shop/item"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetIem(ctx context.Context, request *item.GetItemRequest, callOptions ...callopt.Option) (r *item.GetItemResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kItemServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kItemServiceClient struct {
	*kClient
}

func (p *kItemServiceClient) GetIem(ctx context.Context, request *item.GetItemRequest, callOptions ...callopt.Option) (r *item.GetItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetIem(ctx, request)
}
