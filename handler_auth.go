package yugle

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

type UserUpdate struct {
	UserId     uint   `json:"userId" required`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Gender     int8   `json:"gender"`
	Avatar     string `json:"avatar"`
	BirthYear  int    `json:"birthYear"`
	BirthMonth int8   `json:"birthMonth"`
	BirthDay   int8   `json:"birthDay"`
}

func GetUserPageHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	authUserQuery := &AuthUserQuery{}
	c.ShouldBindJSON(authUserQuery)
	c.JSON(http.StatusOK, OkData(GetAuthUserPage(pageParams.Page, pageParams.Size, authUserQuery)))
}

func GetCurrentUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	username, ok := claims["id"].(string)
	if !ok {
		c.JSON(http.StatusOK, Fail())
		return
	}
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
	authUser := &AuthUser{}
	authUser.Username = userCreate.Username
	authUser.Enabled = userCreate.Enabled
	SaveAuthUser(authUser)
	c.JSON(http.StatusOK, Ok())
}

func UpdateUserHandler(c *gin.Context) {
	userUpdate := &UserUpdate{}
	c.ShouldBindJSON(userUpdate)
	user := GetUserById(userUpdate.UserId)
	user.UserId = userUpdate.UserId
	user.Phone = userUpdate.Phone
	user.Email = userUpdate.Email
	user.Gender = userUpdate.Gender
	user.Avatar = userUpdate.Avatar
	user.BirthYear = userUpdate.BirthYear
	user.BirthMonth = userUpdate.BirthMonth
	user.BirthDay = userUpdate.BirthDay
	SaveUser(user)
	c.JSON(http.StatusOK, Ok())
}

func ChangeUserStatusHandler(c *gin.Context) {
	userId, userIdValid := strconv.ParseUint(c.Param("userId"), 10, 32)
	enabled, enabledValid := strconv.ParseBool(c.Query("enabled"))
	if userIdValid != nil || enabledValid != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	user := GetAuthUserById(uint(userId))
	user.Enabled = enabled
	SaveAuthUser(user)
	c.JSON(http.StatusOK, Ok())
}

func ChangePasswordHandler(c *gin.Context) {

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
