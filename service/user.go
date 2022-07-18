package service

import (
	"douyin/dal/db"
	"douyin/model"
	"errors"
	"strconv"
)

// Register 用户注册
func Register(name, password string) (User, error) {
	if _, err := db.GetUserByName(name); err == nil {
		return User{}, errors.New("账号已存在")
	} else {
		newUser := model.User{
			Name:     name,
			Password: password,
		}
		err := db.AddUser(newUser)
		if err != nil {
			return User{}, errors.New("注册失败")
		}
		t, err := db.GetUserByName(name)
		return User{
			Id:            t.Id,
			Name:          t.Name,
			FollowCount:   t.FollowerCount,
			FollowerCount: t.FollowCount,
			IsFollow:      t.IsFollow,
		}, nil
	}
}

// Login 用户登录
func Login(name, password string) (User, error) {
	if user, err := db.GetUserByName(name); err == nil {
		if user.Password == password {
			return User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowerCount,
				FollowerCount: user.FollowCount,
				IsFollow:      user.IsFollow,
			}, nil
		}
		return User{}, errors.New("密码错误")
	}
	return User{}, errors.New("账号不存在")
}

// UserInfo 用户信息
func UserInfo(userId string) (User, error) {
	id, _ := strconv.Atoi(userId)
	user, err := db.GetUserById(int64(id))
	if err != nil {
		return User{}, err
	} else {
		return User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowerCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}, nil
	}
}
