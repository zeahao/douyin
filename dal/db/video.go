package db

import (
	"douyin/model"
)

// GetFeedList 获取视频feed流
func GetFeedList(latestTime int64) (videos []model.Video, nextTime int64) {
	nextTime = latestTime
	_ = db.Table("video").Where("create_time<?", nextTime).
		Limit(30).Order("create_time desc").
		Find(&videos).Error

	// 更新时间戳
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreateTime
	}

	return videos, nextTime
}

// GetVideoListByAuthor 发表列表
func GetVideoListByAuthor(userId int64) (videos []model.Video) {
	db.Table("video").Where("author_id = ?", userId).Find(&videos)
	return videos
}

// GetVideoById 获取视频
func GetVideoById(videoId int64) (video model.Video) {
	db.Table("video").Where("id = ?", videoId).Take(&video)
	return video
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

// GetVideoListById 批量查找视频
func GetVideoListById(videoId []int64) (videos []model.Video) {
	db.Table("video").Where("id in ?", videoId).Find(&videos)
	return videos
}
