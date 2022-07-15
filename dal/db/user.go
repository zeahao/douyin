package db

import (
	"douyin/config"
	"douyin/model"
)

// SelectUserByName 以姓名查询
func SelectUserByName(name string) (user model.User, err error) {
	err = config.Db.Table("user").Where("name=?", name).Take(&user).Error
	return user, err
}

// SelectUserById 以id查询
func SelectUserById(id int64) (user model.User, err error) {
	err = config.Db.Table("user").Where("id=?", id).Take(&user).Error
	return user, err
}

// InsertUser 添加用户
func InsertUser(user model.User) (err error) {
	err = config.Db.Table("user").Create(&user).Error
	return err
}

// UpdateUser 修改用户数据
func UpdateUser(user model.User) (err error) {
	err = config.Db.Table("user").Save(user).Error
	return err
}

// SelectUserList 批量查询
func SelectUserList(id []int64) (users []model.User, err error) {
	err = config.Db.Table("user").Where("id in ?", id).Find(&users).Error
	return users, err
}
