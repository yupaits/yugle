package yugle

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

var authWare *jwt.GinJWTMiddleware

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

	authWare = AuthWare()

	//设置静态资源路径，html模板路径
	router.Static("/static", appConfig.StaticFolder)
	router.StaticFile("/favicon.ico", appConfig.FaviconPath)
	router.LoadHTMLGlob(appConfig.HtmlTemplateDir)

	//404路由
	router.NoRoute(NoRoute)

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
		api.GET("/picture/bing", GetBingPicturesHandler)

		api.GET("/picture/shot_on_oneplus", GetShotOnOnePlusPicturesHandler)

		api.GET("/task/page", GetTasksHandler)

		api.POST("/user/page", func(c *gin.Context) {
			authorize("action:user:page", c)
		}, GetUserPageHandler)

		api.GET("/user/current", GetCurrentUser)

		api.POST("/user", func(c *gin.Context) {
			authorize("action:user:create", c)
		}, AddUserHandler)

		api.PUT("/user/:userId", UpdateUserHandler)

		api.PUT("/user/:userId/status", func(c *gin.Context) {
			authorize("action:user:status:update", c)
		}, ChangeUserStatusHandler)

		api.POST("/user/password", ModifyPasswordHandler)

		api.GET("/user_roles/:userId", func(c *gin.Context) {
			authorize("action:user:role:list", c)
		}, GetUserRolesByUserIdHandler)

		api.POST("/user/roles/assign", func(c *gin.Context) {
			authorize("action:user:role:assign", c)
		}, AssignRolesHandler)

		api.GET("/role/page", func(c *gin.Context) {
			authorize("action:role:page", c)
		}, GetRolePageHandler)

		api.GET("/roles", func(c *gin.Context) {
			authorize("action:role:list", c)
		}, ListRoleHandler)

		api.POST("/role", func(c *gin.Context) {
			authorize("action:role:create", c)
		}, AddRoleHandler)

		api.PUT("/role/:roleId", func(c *gin.Context) {
			authorize("action:role:update", c)
		}, UpdateRoleHandler)

		api.DELETE("/role/:roleId", func(c *gin.Context) {
			authorize("action:role:delete", c)
		}, DeleteRoleHandler)

		api.GET("/role_permissions/:roleId", func(c *gin.Context) {
			authorize("action:role:permission:list", c)
		}, GetRolePermissionsByRoleIdHandler)

		api.POST("/role/permissions/assign", func(c *gin.Context) {
			authorize("action:role:permission:assign", c)
		}, AssignPermissionsHandler)

		api.POST("/permission/page", func(c *gin.Context) {
			authorize("action:permission:page", c)
		}, GetPermissionPageHandler)

		api.GET("/permissions", func(c *gin.Context) {
			authorize("action:permission:list", c)
		}, ListPermissionHandler)

		api.POST("/permission", func(c *gin.Context) {
			authorize("action:permission:create", c)
		}, AddPermissionHandler)

		api.PUT("/permission/:permissionId", func(c *gin.Context) {
			authorize("action:permission:update", c)
		}, UpdatePermissionHandler)

		api.DELETE("/permission/:permissionId", func(c *gin.Context) {
			authorize("action:permission:delete", c)
		}, DeletePermissionHandler)
	}

	//启动定时任务
	StartSpiderScheduler()

	log.Println("Server started at:", "http://localhost:"+appConfig.Port+"/index")
	log.Fatal(router.Run(":" + appConfig.Port))
}
