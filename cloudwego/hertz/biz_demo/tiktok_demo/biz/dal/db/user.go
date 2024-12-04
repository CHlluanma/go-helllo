package db

import (
	"github.com/chhz0/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/pkg/constants"
	"github.com/chhz0/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/pkg/errno"
)

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

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// QueryUser query user info
func QueryUser(username string) (*User, error) {
	var user User
	if err := DB.Where("user_name = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// QueryUserById get user in the database by user id
func QueryUserById(user_id int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return nil, err
	}
	if user == (User{}) {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return &user, nil
}

// VerfiyUser verify username and password in the db
func VerfiyUser(userName, password string) (int64, error) {
	var user User
	if err := DB.Where("user_name = ? AND password = ?", userName, password).Find(&user).Error; err != nil {
		return 0, err
	}
	if user.ID == 0 {
		err := errno.PasswordIsNotVerfied
		return user.ID, err
	}
	return user.ID, nil
}

// CheckUserExistById check whether the user exist in the db
func CheckUserExistById(user_id int64) (bool, error) {
	var user User
	if err := DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
