package model

// Favorite 点赞
type Favorite struct {
	UserId  int64 `json:"user_id"`  //点赞用户ID
	VideoId int64 `json:"video_id"` //被点赞视频ID
}
