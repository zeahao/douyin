package service

import (
	"douyin/dal/db"
	"douyin/model"
	"time"
)

// CommentAction 评论
func CommentAction(userId int64, videoId int64, content string) (err error) {
	err = db.AddComment(model.Comment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    content,
		CreateDate: time.Now().Format("01-02"),
	})
	return err
}

// GetCommentId 查询评论id
func GetCommentId(userId, videoId int64) (comment Comment) {
	t, _ := db.GetComment(userId, videoId)
	comment.Id = t.Id
	comment.CreateDate = t.CreateDate
	return comment
}

// DelComment 删除评论
func DelComment(commentId int64) (err error) {
	err = db.DelComment(commentId)
	return err
}
