package db

import (
	"douyin/model"
)

// GetFeedList 获取视频feed流
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

// GetVideoListByAuthor 发表列表
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
