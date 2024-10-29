package dal

import "github.com/ahang7/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/pkg/constants"

type User struct {
	ID              int64  `json:"id,omitempty"`
	UserName        string `json:"user_name,omitempty"`
	Password        string `json:"password,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
}

func (User) TableName() string {
	return constants.UserTableName
}
