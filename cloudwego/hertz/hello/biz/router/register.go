// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	hello_world "github.com/CHlluanma/go-hello/cloudwego/hertz/hello/biz/router/hello/world"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	hello_world.Register(r)

}
