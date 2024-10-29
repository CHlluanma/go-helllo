package mysql

import (
	"github.com/ahang7/go-hello/cloudwego/hertz/biz_demo/hertz_gorm/biz/model"
)

func CreateUser(users []*model.User) error {
	return DB.Create(users).Error
}

func DeleteUser(userId int64) error {
	return DB.Where("id = ?", userId).Delete(&model.User{}).Error
}

func UpdateUser(user *model.User) error {
	return DB.Updates(user).Error
}

func QueryUser(keyword *string, page, pageSize int64) ([]*model.User, int64, error) {
	db := DB.Model(model.User{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(DB.Or("name LIKE ?", "%"+*keyword+"%")).
			Or("introduce LIKE ?", "%"+*keyword+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var users []*model.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
