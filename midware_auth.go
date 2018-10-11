package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

/*func AuthWare() *jwt.GinJWTMiddleware {
	authWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "yugle",
		Key:        []byte("yupaits"),
		Timeout:    10 * time.Second,
		MaxRefresh: 30 * time.Second,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			return authenticate(userID, password)
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Unauthorized: func(c *gin.Context, status int, message string) {
			c.JSON(status, CodeFail(UNAUTHORIZED))
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authWare
}

//用户认证，检查用户名密码是否匹配
func authenticate(username string, password string) (interface{}, error) {
	user := GetAuthUser(username)
	if user.ID != 0 && user.Password == password && user.Enabled {
		return user, nil
	}
	return nil, jwt.ErrFailedAuthentication
}*/

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
