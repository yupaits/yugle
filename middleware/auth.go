package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"time"
	"yugle/common/result"
	"yugle/model"
)

func AuthWare() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "yugle",
		Key:        []byte("yupaits"),
		Timeout:    1 * time.Hour,
		MaxRefresh: 47 * time.Hour,
		Authenticator: func(userID string, password string, c *gin.Context) (interface{}, bool) {
			return authenticate(userID, password)
		},
		Authorizator: func(username interface{}, c *gin.Context) bool {
			return authorize(username, c.Request.Method, c.Request.URL.Path)
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Unauthorized: func(c *gin.Context, status int, message string) {
			c.JSON(status, result.CodeFail(result.UNAUTHORIZED))
		},
	}
}

//用户认证，检查用户名密码是否匹配
func authenticate(username string, password string) (interface{}, bool) {
	user := model.GetAuthUser(username)
	if &user != nil && user.Password == password {
		return user, true
	}
	return nil, false
}

//用户授权
func authorize(username interface{}, method string, path string) bool {
	if name, ok := username.(string); ok {
		user := model.GetAuthUser(name)
		if &user != nil {
			return true
		}
	}
	return false
}
