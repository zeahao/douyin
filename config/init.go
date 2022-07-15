package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	user := "root"
	password := "94364426"
	url := "127.0.0.1:3306"
	table := "douyin"
	Db, _ = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@(%v)/%v", user, password, url, table)))
	if err := Db.Error; err != nil {
		fmt.Println("链接失败：", err)
		return
	}
	//tt.URL = tt.GetURL()
	//tt.UserIdSequence, _ = GetLastId()
}
