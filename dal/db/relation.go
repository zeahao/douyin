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
