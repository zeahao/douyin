package db

import (
	"douyin/model"
)

// GetFeedList 获取视频feed流
func GetFeedList(latestTime int64, userId int64) (videos []model.Video, nextTime int64) {
	nextTime = latestTime
	_ = db.Table("video").Where("create_time<?", nextTime).
		Limit(30).Order("create_time desc").
		Find(&videos).Error

	// 更新时间戳
	nextTime = videos[len(videos)-1].CreateTime

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

// GetVideoById 获取视频
func GetVideoById(videoId int64) (video model.Video, err error) {
	err = db.Table("video").Where("id = ?", videoId).Take(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

// UpdateVideo 修改视频数据
func UpdateVideo(video model.Video) (err error) {
	err = db.Table("video").Save(&video).Error
	if err != nil {
		return err
	}
	return nil
}

// AddVideo 添加视频
func AddVideo(newVideo model.Video) error {
	err := db.Table("video").Create(&newVideo).Error
	return err
}
