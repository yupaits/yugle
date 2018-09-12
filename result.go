package yugle

import "github.com/gin-gonic/gin"

const (
	OK               = 200
	FAIL             = 201
	LoginFail        = 202
	ParamsError      = 203
	DataNotFound     = 204
	DataCannotDelete = 205
	DataValidError   = 206
	DatabaseError    = 207
	UNAUTHORIZED     = 401
	FORBIDDEN        = 403
	NotFound         = 404
)

var messages = map[int]string{
	OK:               "成功",
	FAIL:             "失败",
	LoginFail:        "用户名或密码错误",
	ParamsError:      "传入参数有误",
	DataNotFound:     "数据不存在或已被删除",
	DataCannotDelete: "数据无法删除",
	DataValidError:   "查找的数据校验出错",
	DatabaseError:    "数据库错误",
	UNAUTHORIZED:     "身份认证失败",
	FORBIDDEN:        "权限不够，禁止访问",
	NotFound:         "文件或目录未找到",
}

func Ok() gin.H {
	return gin.H{"code": OK, "msg": messages[OK]}
}

func OkData(data interface{}) gin.H {
	return gin.H{"code": OK, "msg": messages[OK], "data": data}
}

func Fail() gin.H {
	return gin.H{"code": FAIL, "msg": messages[FAIL]}
}

func CodeFail(code int) gin.H {
	return gin.H{"code": code, "msg": messages[code]}
}
