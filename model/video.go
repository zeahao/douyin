package model

// Video 视频
type Video struct {
	Id            int64  `json:"id,omitempty"`                       //视频ID
	AuthorId      int64  `json:"author_id"`                          //作者ID
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"` //视频链接
	CoverUrl      string `json:"cover_url,omitempty"`                //视频封面图链接
	FavoriteCount int64  `json:"favorite_count,omitempty"`           //视频点赞数
	CommentCount  int64  `json:"comment_count,omitempty"`            //视频评论数
	IsFavorite    bool   `json:"is_favorite,omitempty"`              //视频点赞数是否点赞
	Title         string `json:"title,omitempty"`                    //视频描述
	CreateTime    int64  `json:"create_time"`                        //视频上传时间
}
