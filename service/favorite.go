package service

import (
	"douyin/dal/db"
	"douyin/model"
)

// FavoriteAction 视频点赞
func FavoriteAction(userId, videoId int64) {
	_ = db.AddFavorite(model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	})
}

// DelFavorite 取消点赞
func DelFavorite(userId, videoId int64) {
	_ = db.DelFavorite(userId, videoId)
}
