package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yugle/appconfig"
	"yugle/handler"
)

var authMiddleware *jwt.GinJWTMiddleware

func main() {
	router := gin.Default()

	config := appconfig.DefaultConfig()

	router.Static("/static", config.StaticFolder)
	router.StaticFile("/favicon.ico", config.FaviconPath)
	router.LoadHTMLGlob(config.HtmlTemplateDir)

	router.NoRoute(handler.NoRoute)

	//页面路由
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//API路由
	router.POST("/login", authMiddleware.LoginHandler)

	router.Run(":" + config.Port)
}
