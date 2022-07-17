package service

import (
	"douyin/dal/db"
	"douyin/model"
	"sync"
)

// FavoriteAction 视频点赞
func FavoriteAction(userId, videoId int64) {
	_ = db.AddFavorite(model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	})
	video, err := db.GetVideoById(videoId)
	if err != nil {
		return
	}
	video.FavoriteCount++
	err = db.UpdateVideo(video)
	if err != nil {
		return
	}

}

// DelFavorite 取消点赞
func DelFavorite(userId, videoId int64) {
	_ = db.DelFavorite(userId, videoId)
	video, err := db.GetVideoById(videoId)
	if err != nil {
		return
	}
	video.FavoriteCount--
	err = db.UpdateVideo(video)
	if err != nil {
		return
	}
}

// GetFavoriteList 获取点赞视频列表
func GetFavoriteList(userId int64) (videos []Video) {
	l := db.GetFavoriteVideoIdList(userId)
	list, _ := db.GetVideoListById(l)
	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)
		go func(v model.Video) {
			defer wg.Done()
			user, _ := db.GetUserById(v.AuthorId)
			videos = append(videos, Video{
				Id: v.Id,
				Author: User{
					Id:            user.Id,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					IsFollow:      user.IsFollow,
				},
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				IsFavorite:    v.IsFavorite,
			})
		}(v)
	}
	wg.Wait()

	return videos
}