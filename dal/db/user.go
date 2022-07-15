package db

import (
	"douyin/model"
)

// SelectUserByName 以姓名查询
func SelectUserByName(name string) (user model.User, err error) {
	err = db.Table("user").Where("name=?", name).Take(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// SelectUserById 以id查询
func SelectUserById(id int64) (user model.User, err error) {
	err = db.Table("user").Where("id=?", id).Take(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// InsertUser 添加用户
func InsertUser(user model.User) (err error) {
	err = db.Table("user").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser 修改用户数据
func UpdateUser(user model.User) (err error) {
	err = db.Table("user").Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// SelectUserList 批量查询
func SelectUserList(id []int64) (users []model.User, err error) {
	err = db.Table("user").Where("id in ?", id).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
