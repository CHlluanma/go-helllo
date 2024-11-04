package main

import (
	"context"

	api "github.com/CHlluanma/go-hello/cloudwego/kitex/hello/kitex_gen/api"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, request *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{Message: request.Message}

	return
}

// Add implements the HelloImpl interface.
func (s *HelloImpl) Add(ctx context.Context, request *api.AddRequest) (resp *api.AddResponse, err error) {
	resp = &api.AddResponse{Sum: request.First + request.Second}
	return
}
