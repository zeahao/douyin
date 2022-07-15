package model

// Comment 评论
type Comment struct {
	Id         int64  `json:"id,omitempty"`          //评论ID
	UserId     int64  `json:"user_id,omitempty"`     // 作者ID
	VideoId    int64  `json:"video_id,omitempty"`    // 评论的视频ID
	Content    string `json:"content,omitempty"`     //评论内容
	CreateDate string `json:"create_date,omitempty"` //评论时间
}
