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
