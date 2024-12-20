// Code generated by hertz generator.

package user

import (
	"context"

	user "github.com/chhz0/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/biz/model/basic/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// User .
// @router douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DouyinUserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}
