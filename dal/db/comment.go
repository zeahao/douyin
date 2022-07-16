package db

import "douyin/model"

// AddComment 添加评论
func AddComment(comment model.Comment) (err error) {
	err = db.Table("comment").Create(&comment).Error
	return err
}

// GetComment 查询评论
func GetComment(userId, videoId int64) (comment model.Comment, err error) {
	err = db.Table("comment").
		Where("user_id = ? and video_id = ?", userId, videoId).
		Take(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// DelComment 删除评论
func DelComment(commentId int64) (err error) {
	err = db.Table("comment").Where("id=?", commentId).Delete(&model.Comment{}).Error
	return err
}
