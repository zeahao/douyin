package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []service.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment service.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			videoId, _ := strconv.Atoi(c.Query("video_id"))
			err := service.CommentAction(user.Id, int64(videoId), text)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{
						StatusCode: 1,
						StatusMsg:  err.Error(),
					},
				})
				return
			}
			comment := service.GetCommentId(user.Id, int64(videoId))
			comment.User = user
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "评论成功"},
				Comment: comment})
			return
		} else if actionType == "2" {
			commentId, _ := strconv.Atoi(c.Query("comment_id"))
			service.DelComment(int64(commentId))
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "删除成功"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	comments := service.GetCommentList(int64(videoId))

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
