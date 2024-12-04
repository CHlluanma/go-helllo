// Code generated by hertz generator.

package publish

import (
	"context"

	publish "github.com/chhz0/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/biz/model/basic/publish"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.DouyinPublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.DouyinPublishListResponse)

	c.JSON(consts.StatusOK, resp)
}
