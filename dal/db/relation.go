package db

import "douyin/model"

// AddRelation 添加关注
func AddRelation(relation model.Relation) (err error) {
	err = db.Table("relation").Create(&relation).Error
	return err
}

// DelRelation 删除关注
func DelRelation(relation model.Relation) (cnt int64) {
	db.Table("relation").Where("user_id=? and to_user_id=?", relation.UserId, relation.ToUserId).Count(&cnt).Delete(&relation)
	return cnt
}

// GetToUserIdList 获取关注列表
func GetToUserIdList(userId int64) (toUserIdList []int64) {
	db.Table("relation").Where("user_id=?", userId).
		Select("to_user_id").Find(&toUserIdList)
	return toUserIdList
}

// GetUserIdList 获取粉丝列表
func GetUserIdList(userId int64) (toUserIdList []int64) {
	db.Table("relation").Where("to_user_id=?", userId).
		Select("user_id").Find(&toUserIdList)
	return toUserIdList
}

// IsRelation 判断是否关注
func IsRelation(userId, toUserId int64) bool {
	err := db.Table("relation").Where("user_id=? and to_user_id=?", userId, toUserId).Error
	if err != nil {
		return false
	}
	return true
}
