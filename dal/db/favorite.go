package db

import "douyin/model"

// AddFavorite 添加点赞
func AddFavorite(favorite model.Favorite) (err error) {
	err = db.Table("favorite").Create(&favorite).Error
	return err
}

// DelFavorite 删除点赞
func DelFavorite(userId, videoId int64) (err error) {
	err = db.Table("favorite").Where("user_id=? and video_id = ?", userId, videoId).Delete(&model.Favorite{}).Error
	return err
}

// GetFavoriteVideoIdList 批量获取点赞视频id
func GetFavoriteVideoIdList(userId int64) (videoIdList []int64) {
	db.Table("favorite").Where("user_id=?", userId).Select("video_id").Find(&videoIdList)
	return videoIdList
}
