package yugle

import (
	"github.com/gin-contrib/authz"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
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

	//设置鉴权配置
	enforcer := GetEnforcer()
	router.Use(authz.NewAuthorizer(enforcer))

	//设置session相关配置
	store := cookie.NewStore([]byte("yupaits"))
	router.Use(sessions.Sessions("yugle.session", store))

	//设置静态资源路径，html模板路径
	router.Static("/static", appConfig.StaticFolder)
	router.StaticFile("/favicon.ico", appConfig.FaviconPath)
	router.LoadHTMLGlob(appConfig.HtmlTemplateDir)

	//404路由
	router.NoRoute(NoRoute)

	//页面路由
	router.GET("/", func(c *gin.Context) {
		routePage("index.html", c)
	})
	router.GET("/index", func(c *gin.Context) {
		routePage("index.html", c)
	})
	router.GET("/login", func(c *gin.Context) {
		routePage("login.html", c)
	})

	//认证授权路由
	auth := router.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/refresh_token")
	}

	//API路由
	api := router.Group("/api")
	{
		api.GET("/picture/bing", GetBingPicturesHandler)
		api.GET("/picture/shot_on_oneplus", GetShotOnOnePlusPicturesHandler)

		api.GET("/task/page", GetTasksHandler)

		api.POST("/user/page", GetUserPageHandler)
		api.GET("/user/current", GetCurrentUser)
		api.POST("/user", AddUserHandler)
		api.PUT("/user/:userId", UpdateUserHandler)
		api.PUT("/user/:userId/status", ChangeUserStatusHandler)
		api.POST("/user/password", ModifyPasswordHandler)
		api.GET("/user_roles/:userId", GetUserRolesByUserIdHandler)
		api.POST("/user/roles/assign", AssignRolesHandler)

		api.GET("/role/page", GetRolePageHandler)
		api.GET("/roles", ListRoleHandler)
		api.POST("/role", AddRoleHandler)
		api.PUT("/role/:roleId", UpdateRoleHandler)
		api.DELETE("/role/:roleId", DeleteRoleHandler)
		api.GET("/role_permissions/:roleId", GetRolePermissionsByRoleIdHandler)
		api.POST("/role/permissions/assign", AssignPermissionsHandler)

		api.POST("/permission/page", GetPermissionPageHandler)
		api.GET("/permissions", ListPermissionHandler)
		api.POST("/permission", AddPermissionHandler)
		api.PUT("/permission/:permissionId", UpdatePermissionHandler)
		api.DELETE("/permission/:permissionId", DeletePermissionHandler)
	}

	//启动定时任务
	StartSpiderScheduler()

	log.Println("Server started at:", "http://localhost:"+appConfig.Port+"/index")
	log.Fatal(router.Run(":" + appConfig.Port))
}
