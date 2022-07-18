package service

import (
	"douyin/dal/db"
	"douyin/model"
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
func DelRelation(userId, toUserId int64) (err error) {
	err = db.DelRelation(model.Relation{
		UserId:   userId,
		ToUserId: toUserId,
	})

	if err != nil {
		return err
	}
	user, _ := db.GetUserById(userId)
	user.FollowCount--
	_ = db.UpdateUser(user)

	user, _ = db.GetUserById(toUserId)
	user.FollowerCount--
	_ = db.UpdateUser(user)
	return nil
}
