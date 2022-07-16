package main

import (
	"douyin/config"
	"douyin/dal/db"
	"douyin/router"
	"github.com/gin-gonic/gin"
)

func init() {
	user := "root"
	password := "94364426"
	url := "127.0.0.1:3306"
	database := "douyin"
	db.InitMysql(user, password, url, database)
	config.InitGetURL()
}

func main() {
	r := gin.Default()

	router.InitRouter(r)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
