package service

import (
	"douyin/config"
	"douyin/dal/db"
	"douyin/model"
	"douyin/util"
	"sync"
	"time"
)

// GetFeedList 获取视频feed流
func GetFeedList(latestTime int64, userId int64) (videos []Video, nextTime int64) {
	list, nextTime := db.GetFeedList(latestTime, userId)

	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)
		go func(v model.Video) {
			defer wg.Done()
			user, _ := db.GetUserById(v.AuthorId)
			//此视频是否已被点赞(好像多余了)
			//favorite := 0
			//db.Debug().Table("favorite").Select("user_id").Where("user_id=? and video_id=?", userId, v.Id).Take(&favorite)
			//if favorite > 0 {
			//	v.IsFavorite = true
			//} else {
			//	v.IsFavorite = false
			//}
			//follow := 0
			//db.Debug().Table("relation").Select("user_id").Where("user_id=? and to_user_id=?", userId, v.AuthorId).Take(&follow)
			//if follow > 0 || userId == user.Id {
			//	user.IsFollow = true
			//} else {
			//	user.IsFollow = false
			//}
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
	return videos, nextTime
}

// GetPublishList 获取发布列表
func GetPublishList(userId int64) (videos []Video) {
	list, _ := db.GetVideoListByAuthor(userId)

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

// Publish 视频发布接口
func Publish(user User, videoName string, title string) (err error) {
	// 获取视频封面图
	imgName := util.GetImage(videoName)

	err = db.AddVideo(model.Video{
		AuthorId:   user.Id,
		PlayUrl:    config.URL + ":8080/Video/" + videoName,
		CoverUrl:   config.URL + ":8080/Image/" + imgName + ".jpeg",
		Title:      title,
		CreateTime: time.Now().UnixMilli(),
	})
	return err
}
