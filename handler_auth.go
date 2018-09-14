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

type UserCreate struct {
	Username string `json:"username" required`
	Enabled  bool   `json:"enabled" required`
}

func GetUserPageHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	authUserQuery := &AuthUserQuery{}
	c.ShouldBindJSON(authUserQuery)
	c.JSON(http.StatusOK, OkData(GetAuthUserPage(pageParams.Page, pageParams.Size, authUserQuery)))
}

func GetUserByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	authUser := GetAuthUser(username)
	if authUser == nil {
		c.JSON(http.StatusOK, CodeFail(DataNotFound))
		return
	}
	userVO := &UserVO{}
	userVO.Username = authUser.Username
	userVO.Enabled = authUser.Enabled
	userVO.User = *GetUserById(authUser.ID)
	userVO.User.UserId = authUser.ID
	c.JSON(http.StatusOK, OkData(userVO))
}

func AddUserHandler(c *gin.Context) {
	userCreate := &UserCreate{}
	c.ShouldBindJSON(userCreate)
}

func UpdateUserHandler(c *gin.Context) {

}

func ChangeUserStatusHandler(c *gin.Context) {

}

func UpdateUserAuthHandler(c *gin.Context) {

}

func GetUserRolesByUserIdHandler(c *gin.Context) {

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

func GetRolePermissionsByRoleIdHandler(c *gin.Context) {

}

func GetPermissionPageHandler(c *gin.Context) {

}

func ListPermissionHandler(c *gin.Context) {

}

func AddPermissionHandler(c *gin.Context) {

}

func UpdatePermissionHandler(c *gin.Context) {

}
