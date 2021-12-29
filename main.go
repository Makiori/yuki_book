package main

import (
	"fmt"
	"net/http"
	"time"
	"yuki_book/model"
	"yuki_book/router"
	"yuki_book/util/conf"
	"yuki_book/util/gredis"
	"yuki_book/util/logging"
	"yuki_book/util/times"
	"yuki_book/util/validator"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logging.Setup()
	conf.Setup()
	gredis.Setup()
	validator.Setup()
	model.Init()
}

// @title yuki_book
// @version 1.0
// @description 借书系统
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logging.Info("设置服务器时区为东八区")
	timeLocal := time.FixedZone("CST", 8*3600)
	time.Local = timeLocal
	logging.Info("当前时间:", times.ToStr())
	gin.SetMode(conf.Data.Server.RunMode)
	httpPort := fmt.Sprintf(":%d", conf.Data.Server.HttpPort)
	server := &http.Server{
		Addr:           httpPort,
		Handler:        router.InitRouter(),
		ReadTimeout:    conf.Data.Server.ReadTimeout,
		WriteTimeout:   conf.Data.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	logging.Info("开始监听服务", httpPort)
	if err := server.ListenAndServe(); err != nil {
		logging.Fatal("启动服务失败", err)
	}
}
