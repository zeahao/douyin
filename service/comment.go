package service

import (
	"douyin/dal/db"
	"douyin/model"
	"sync"
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

// GetCommentId 查询评论
func GetCommentId(userId, videoId int64) (comment Comment) {
	t, _ := db.GetComment(userId, videoId)
	comment.Id = t.Id
	comment.CreateDate = t.CreateDate
	return comment
}

// DelComment 删除评论
func DelComment(commentId int64) {
	db.DelComment(commentId)
}

// GetCommentList 获取评论列表
func GetCommentList(videoId int64) (comments []Comment) {
	t, _ := db.GetCommentList(videoId)

	// 更新视频评论数
	video, _ := db.GetVideoById(videoId)
	video.CommentCount = int64(len(t))
	_ = db.UpdateVideo(video)

	wg := sync.WaitGroup{}
	for _, c := range t {
		wg.Add(1)
		go func(c model.Comment) {
			defer wg.Done()
			user, _ := db.GetUserById(c.UserId)
			comments = append(comments, Comment{
				Id: c.Id,
				User: User{
					Id:            user.Id,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					IsFollow:      user.IsFollow,
				},
				Content:    c.Content,
				CreateDate: c.CreateDate,
			})
		}(c)
	}
	wg.Wait()
	return comments
}
