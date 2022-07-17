package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if _, exist := usersLoginInfo[token]; exist {
		userId, _ := strconv.Atoi(c.Query("user_id"))
		videoId, _ := strconv.Atoi(c.Query("video_id"))
		if actionType == "1" {
			service.FavoriteAction(int64(userId), int64(videoId))
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "点赞成功"})
		} else if actionType == "2" {
			service.DelFavorite(int64(userId), int64(videoId))
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "取消点赞成功"})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "参数有误"})
		}
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	videos := service.GetFavoriteList(int64(userId))
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "查询成功",
		},
		VideoList: videos,
	})
}
