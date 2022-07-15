package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysql(user, password, url, database string) {
	db, _ = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@(%v)/%v", user, password, url, database)))
	if err := db.Error; err != nil {
		fmt.Println("链接失败：", err)
		return
	}
}
