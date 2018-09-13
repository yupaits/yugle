package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthUserQuery struct {
	Keyword string `json:"keyword"`
	Enabled bool   `json:"enabled"`
}

type UserVO struct {
	Username string
	Enabled  bool
	User
}

func GetUserPageHandler(c *gin.Context) {
	var pageParams PageParams
	if c.ShouldBindQuery(&pageParams) != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
	}
	if pageParams.Page == 0 {
		pageParams.Page = DefaultPage
	}
	if pageParams.Size == 0 {
		pageParams.Size = DefaultSize
	}
	authUserQuery := &AuthUserQuery{}
	c.ShouldBindJSON(authUserQuery)
	c.JSON(http.StatusOK, OkData(GetAuthUserPage(pageParams.Page, pageParams.Size, authUserQuery)))
}

func GetUserByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	GetAuthUser(username)
}

func AddUserHandler(c *gin.Context) {

}

func UpdateUserHandler(c *gin.Context) {

}

func ChangeUserStatusHandler(c *gin.Context) {

}

func UpdateUserAuthHandler(c *gin.Context) {

}

func GetRolePageHandler(c *gin.Context) {

}

func ListRoleHandler(c *gin.Context) {

}

func GetRoleByIdHandler(c *gin.Context) {

}

func AddRoleHandler(c *gin.Context) {

}

func UpdateRoleHandler(c *gin.Context) {

}

func GetPermissionPageHandler(c *gin.Context) {

}

func ListPermissionHandler(c *gin.Context) {

}

func AddPermissionHandler(c *gin.Context) {

}

func UpdatePermissionHandler(c *gin.Context) {

}
