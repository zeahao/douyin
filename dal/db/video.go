package db

import (
	"douyin/model"
)

func GetFeedList(latestTime int64, userId int64) (videos []model.Video, nextTime int64) {
	var t model.Video
	nextTime = latestTime
	for i := 0; i < 30; i++ {
		err := db.Table("video").Where("create_time<?", nextTime).Last(&t).Error
		if err != nil {
			break
		}
		videos = append(videos, t)
		//获取本次缓存视频时间戳
		nextTime = t.CreateTime
	}

	//var wg sync.WaitGroup
	//for _, v := range videos {
	//	wg.Add(1)
	//	go func(v model.Video) {
	//		defer wg.Done()
	//		user,_ := GetUserById(v.AuthorId)
	//
	//		//此视频是否已被点赞(好像多余了)
	//		favorite := 0
	//		db.Debug().Table("favorite").Select("user_id").Where("user_id=? and video_id=?", userId, v.Id).Take(&favorite)
	//		if favorite > 0 {
	//			v.IsFavorite = true
	//		} else {
	//			v.IsFavorite = false
	//		}
	//		follow := 0
	//		db.Debug().Table("relation").Select("user_id").Where("user_id=? and to_user_id=?", userId, v.AuthorId).Take(&follow)
	//		if follow > 0 || userId == user.Id {
	//			user.IsFollow = true
	//		} else {
	//			user.IsFollow = false
	//		}
	//	}(v)
	//}
	//wg.Wait()
	return videos, nextTime
}

func GetVideoListByAuthor(userId int64) (videos []model.Video, err error) {
	err = db.Table("video").Where("author_id = ?", userId).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func AddVideo(newVideo model.Video) error {
	err := db.Table("video").Create(&newVideo).Error
	return err
}

func GetPublishList(userId int64) []model.Video {
	user, _ := GetUserById(userId)
	videos, _ := GetVideoListByAuthor(user.Id)
	return videos
}

//func GetVideo(videoId int64) (video tt.Video, err error) {
//	var v model.Video
//	err = Db.Table("video").Where("id = ?", videoId).Take(&v).Error
//	user, _ := GetUser(v.AuthorId)
//	video = tt.Video{
//		Id:            v.Id,
//		Author:        user,
//		PlayUrl:       v.PlayUrl,
//		CoverUrl:      v.CoverUrl,
//		FavoriteCount: v.FavoriteCount,
//		CommentCount:  v.CommentCount,
//		IsFavorite:    v.IsFavorite,
//		Title:         v.Title,
//	}
//	return video, err
//}
//
//func UpdateFavoriteCount(videoId, favoriteCount int64) (err error) {
//	err = Db.Table("video").Where("id=?", videoId).Update("favorite_count", favoriteCount).Error
//	return err
//}
//
//func UpdateCommentCount(videoId, commentCount int64) (err error) {
//	err = Db.Debug().Table("video").Where("id=?", videoId).Update("comment_count", commentCount).Error
//	return err
//}
//
//func GetVideoList(videoIds []int64) (videos []tt.Video, err error) {
//	var video []model.Video
//	err = Db.Table("video").Where("id in ?", videoIds).Find(&video).Error
//	videos = make([]tt.Video, 0, len(video))
//	for _, v := range video {
//		user, _ := GetUser(v.AuthorId)
//		videos = append(videos, tt.Video{
//			Id:            v.Id,
//			Author:        user,
//			PlayUrl:       v.PlayUrl,
//			CoverUrl:      v.CoverUrl,
//			FavoriteCount: v.FavoriteCount,
//			CommentCount:  v.CommentCount,
//			IsFavorite:    true,
//			Title:         v.Title,
//		})
//	}
//	return videos, err
//}
