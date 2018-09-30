package yugle

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthWare() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "yugle",
		Key:        []byte("yupaits"),
		Timeout:    1 * time.Hour,
		MaxRefresh: 48 * time.Hour,
		Authenticator: func(userID string, password string, c *gin.Context) (interface{}, bool) {
			return authenticate(userID, password)
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Unauthorized: func(c *gin.Context, status int, message string) {
			c.JSON(status, CodeFail(UNAUTHORIZED))
		},
	}
}

//用户认证，检查用户名密码是否匹配
func authenticate(username string, password string) (interface{}, bool) {
	user := GetAuthUser(username)
	if user.ID != 0 && user.Password == password && user.Enabled {
		return user, true
	}
	return nil, false
}

//请求鉴权
func authorize(permKey string, c *gin.Context) {
	//TODO 从context中取得username，根据username从数据库或缓存中获取permissions，判断permissions是否能匹配permKey
	if false {
		//鉴权失败处理，返回Forbidden信息
		c.Abort()
		c.JSON(http.StatusForbidden, CodeFail(FORBIDDEN))
	} else {
		c.Next()
	}
}
