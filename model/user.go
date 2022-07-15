package model

// User 用户
type User struct {
	Id            int64  `json:"id,omitempty"`             //用户ID
	Name          string `json:"name,omitempty"`           //用户名
	Password      string `json:"password"`                 //用户密码
	FollowCount   int64  `json:"follow_count,omitempty"`   //关注数
	FollowerCount int64  `json:"follower_count,omitempty"` //粉丝数
	IsFollow      bool   `json:"is_follow,omitempty"`      //是否关注
}
