package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"yugle/config"
	"yugle/dbutils"
	"yugle/handler"
	"yugle/middleware"
	"yugle/model"
	"yugle/result"
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
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//API路由
	router.POST("/login", authWare.LoginHandler)
	api := router.Group("/api", authWare.MiddlewareFunc())
	{
		api.GET("", func(c *gin.Context) {
			db := dbutils.Connect()
			defer db.Close()
			authUser1 := &model.AuthUser{Username: "admin", Password: "123456"}
			authUser2 := &model.AuthUser{Model: gorm.Model{ID: 4}, Username: "user", Password: "123456"}
			db.AutoMigrate(authUser1)
			if db.NewRecord(authUser1) {
				db.Create(authUser1)
			}
			if db.NewRecord(authUser2) {
				db.Create(authUser2)
			}
			c.JSON(http.StatusOK, result.Ok())
		})
	}

	log.Println("Server started at:", "http://localhost:"+appConfig.Port+"/index")
	log.Fatal(router.Run(":" + appConfig.Port))
}
