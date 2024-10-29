package pack

import (
	"github.com/ahang7/go-hello/cloudwego/hertz/biz_demo/hertz_gorm/biz/hertz_gen/user_gorm"
	"github.com/ahang7/go-hello/cloudwego/hertz/biz_demo/hertz_gorm/biz/model"
)

func Users(models []*model.User) []*user_gorm.User {
	users := make([]*user_gorm.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func User(model *model.User) *user_gorm.User {
	if model == nil {
		return nil
	}

	return &user_gorm.User{
		ID:        int64(model.ID),
		Name:      model.Name,
		Age:       model.Age,
		Gender:    user_gorm.Gender(model.Gender),
		Introduce: model.Introduce,
	}
}
