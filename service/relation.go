package service

import (
	"douyin/dal/db"
	"douyin/model"
	"sync"
)

// RelationAction 关注操作
func RelationAction(userId, toUserId int64) (err error) {
	err = db.AddRelation(model.Relation{
		UserId:   userId,
		ToUserId: toUserId,
	})
	if err != nil {
		return err
	}
	user, _ := db.GetUserById(userId)
	user.FollowCount++
	_ = db.UpdateUser(user)

	user, _ = db.GetUserById(toUserId)
	user.FollowerCount++
	_ = db.UpdateUser(user)
	return nil
}

// DelRelation 删除关注
func DelRelation(userId, toUserId int64) {
	cnt := db.DelRelation(model.Relation{
		UserId:   userId,
		ToUserId: toUserId,
	})
	if cnt == 0 {
		return
	}
	user, _ := db.GetUserById(userId)
	user.FollowCount--
	_ = db.UpdateUser(user)

	user, _ = db.GetUserById(toUserId)
	user.FollowerCount--
	_ = db.UpdateUser(user)
	return
}

// GetFollowerList 获取用户粉丝列表
func GetFollowerList(userId int64) (users []User) {
	idList := db.GetUserIdList(userId)
	u := db.GetUserList(idList)
	wg := sync.WaitGroup{}
	for _, user := range u {
		wg.Add(1)
		go func(user model.User) {
			defer wg.Done()
			users = append(users, User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      db.IsRelation(userId, user.Id),
			})
		}(user)
	}
	wg.Wait()

	return users
}

// GetFollowList 获取用户关注列表
func GetFollowList(userId int64) (users []User) {
	idList := db.GetToUserIdList(userId)
	u := db.GetUserList(idList)
	wg := sync.WaitGroup{}
	for _, user := range u {
		wg.Add(1)
		go func(user model.User) {
			defer wg.Done()
			users = append(users, User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      true,
			})
		}(user)
	}
	wg.Wait()

	return users
}
