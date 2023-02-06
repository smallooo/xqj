package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
	"xqj/models"
	"xqj/pkg/app"
	"xqj/pkg/gredis"
	"xqj/pkg/logging"
	"xqj/pkg/setting"
	"xqj/pkg/util"
	"xqj/routers"
)

func main() {

	setting.Setup()
	models.Setup()

	logging.Setup()
	gredis.Setup()
	util.Setup()

	go app.RunTask()

	gin.SetMode(setting.ServerSetting.RunMode)

	router := routers.InitRouter()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	server := endless.NewServer(string("localhost:8000"), router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
