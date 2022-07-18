package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []service.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; exist {
		actionType := c.Query("action_type")
		userId, _ := strconv.Atoi(c.Query("user_id"))
		toUserId, _ := strconv.Atoi(c.Query("to_user_id"))
		if actionType == "1" {
			err := service.RelationAction(int64(userId), int64(toUserId))
			if err != nil {
				c.JSON(http.StatusOK, Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "关注成功",
			})
		} else if actionType == "2" {
			err := service.DelRelation(int64(userId), int64(toUserId))
			if err != nil {
				c.JSON(http.StatusOK, Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "取消关注成功",
			})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "参数错误",
			})
		}

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same followed list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []service.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []service.User{DemoUser},
	})
}
