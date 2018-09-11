package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"yugle/config"
	"yugle/handler"
	"yugle/handler/picture"
	"yugle/middleware"
	"yugle/scheduler"
)

var authWare *jwt.GinJWTMiddleware

func main() {
	//获取默认配置信息
	appConfig := config.DefaultConfig()

	//日志文件
	logFile, _ := os.Create("yugle.log")
	//仅输出日志到文件
	//gin.DefaultWriter = io.MultiWriter(logFile)
	//输出日志到文件和控制台
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	gin.SetMode(appConfig.Mode)

	router := gin.Default()

	authWare = middleware.AuthWare()

	//设置静态资源路径，html模板路径
	router.Static("/static", appConfig.StaticFolder)
	router.StaticFile("/favicon.ico", appConfig.FaviconPath)
	router.LoadHTMLGlob(appConfig.HtmlTemplateDir)

	//404路由
	router.NoRoute(handler.NoRoute)

	//页面路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//认证授权路由
	auth := router.Group("/auth")
	{
		auth.POST("/login", authWare.LoginHandler)
		auth.POST("/refresh_token", authWare.RefreshHandler)
	}

	//API路由
	api := router.Group("/api", authWare.MiddlewareFunc())
	{
		api.GET("/picture/bing", picture_handler.GetBingPicturesHandler)
		api.GET("/picture/shot_on_oneplus", picture_handler.GetShotOnOnePlusPicturesHandler)
	}

	//启动定时任务
	scheduler.StartSpiderScheduler()

	log.Println("Server started at:", "http://localhost:"+appConfig.Port+"/index")
	log.Fatal(router.Run(":" + appConfig.Port))
}
