package model

import (
	"yuki_book/model/database"
	"yuki_book/util/conf"
	"yuki_book/util/logging"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Init() {
	var err error
	logging.Info("连接", conf.Data.Database.Type)
	database.DBCon, err = gorm.Open(conf.Data.Database.Type, conf.Data.Database.Url)
	if err != nil {
		logging.Fatalf("model.Setup err: %v", err)
	}
	// debug模式,显示sql语句
	database.DBCon.LogMode(conf.Data.Server.RunMode == gin.DebugMode)
	// 全局禁用表名复数
	database.DBCon.SingularTable(true)
}
