package service

import (
	"douyin/dal/db"
	"douyin/model"
	"sync"
)

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

func GetPublish(userId int64) (videos []Video) {
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
