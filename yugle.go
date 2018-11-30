package yugle

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	//获取默认配置信息
	appConfig := DefaultConfig()

	//日志文件
	logFile, _ := os.Create("yugle.log")
	//仅输出日志到文件
	//gin.DefaultWriter = io.MultiWriter(logFile)
	//输出日志到文件和控制台
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	gin.SetMode(appConfig.Mode)

	router := gin.Default()

	router.GET("/picture/bing", GetBingPicturesHandler)
	router.GET("/picture/shot_on_oneplus", GetShotOnOnePlusPicturesHandler)

	router.GET("/task/page", GetTasksHandler)

	//启动定时任务
	StartSpiderScheduler()

	log.Println("Server started at port:", appConfig.Port)
	log.Fatal(router.Run(":" + appConfig.Port))
}
