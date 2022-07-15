package model

// Relation 关注
type Relation struct {
	UserId   int64 `json:"user_id"`    //关注用户ID
	ToUserId int64 `json:"to_user_id"` //被关注用户ID
}
