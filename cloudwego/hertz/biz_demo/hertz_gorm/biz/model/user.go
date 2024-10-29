package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name,omitempty" column:"name"`
	Gender    int64  `json:"gender,omitempty" column:"gender"`
	Age       int64  `json:"age,omitempty" column:"age"`
	Introduce string `json:"introduce,omitempty" column:"introduce"`
}

func (u *User) TableName() string {
	return "users"
}
