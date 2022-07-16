package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FeedResponse struct {
	Response
	VideoList []service.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	latestTime, _ := strconv.Atoi(c.Query("latest_time"))

	videos, nextTime := service.GetFeedList(int64(latestTime), usersLoginInfo[token].Id)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "获取成功"},
		VideoList: videos,
		NextTime:  nextTime,
	})
}
