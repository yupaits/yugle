package yugle

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	Username string `json:"username" binding:"required"`
	Enabled  bool   `json:"enabled" binding:"required"`
}

type UserUpdate struct {
	UserId     uint   `json:"userId" binding:"required"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Gender     int8   `json:"gender"`
	Avatar     string `json:"avatar"`
	BirthYear  int    `json:"birthYear"`
	BirthMonth int8   `json:"birthMonth"`
	BirthDay   int8   `json:"birthDay"`
}

type PasswordModify struct {
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type RoleVO struct {
	ID          uint
	Key         string
	Name        string
	Description string
	CreatedAt   time.Time
}

type RoleCreate struct {
	Key         string `json:"key" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RoleUpdate struct {
	ID          uint   `json:"id" binding:"required"`
	Key         string `json:"key" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type PermissionQuery struct {
	Keyword  string `json:"keyword"`
	PermType int8   `json:"permType"`
}

type PermissionVO struct {
	ID          uint
	Key         string
	Name        string
	PermType    int8
	Description string
	CreatedAt   time.Time
}

type PermissionCreate struct {
	Key         string `json:"key" binding:"required"`
	Name        string `json:"name" binding:"required"`
	PermType    int8   `json:"permType" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type PermissionUpdate struct {
	ID          uint   `json:"id" binding:"required"`
	Key         string `json:"key" binding:"required"`
	Name        string `json:"name" binding:"required"`
	PermType    int8   `json:"permType" binding:"required"`
	Description string `json:"description" binding:"required"`
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
		c.JSON(http.StatusOK, CodeFail(TokenInfoError))
		return
	}
	authUser := GetAuthUser(username)
	if authUser.ID == 0 {
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
	if userCreate.Username == "" {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	if GetAuthUser(userCreate.Username).ID != 0 {
		c.JSON(http.StatusOK, MsgFail("该用户名已存在"))
		return
	}
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
	userId, userIdErr := strconv.ParseUint(c.Param("userId"), 10, 32)
	enabled, enabledErr := strconv.ParseBool(c.Query("enabled"))
	if userIdErr != nil || enabledErr != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	user := GetAuthUserById(uint(userId))
	user.Enabled = enabled
	SaveAuthUser(user)
	c.JSON(http.StatusOK, Ok())
}

func ModifyPasswordHandler(c *gin.Context) {
	passwordModify := &PasswordModify{}
	c.ShouldBindJSON(passwordModify)
	if passwordModify.OldPassword == "" || passwordModify.NewPassword == "" || passwordModify.ConfirmPassword == "" {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	if passwordModify.OldPassword == passwordModify.NewPassword {
		c.JSON(http.StatusOK, MsgFail("新密码不能与当前密码相同"))
		return
	}
	if passwordModify.NewPassword != passwordModify.ConfirmPassword {
		c.JSON(http.StatusOK, MsgFail("两次输入的新密码不相符"))
		return
	}
	claims := jwt.ExtractClaims(c)
	username, ok := claims["id"].(string)
	if !ok {
		c.JSON(http.StatusOK, CodeFail(TokenInfoError))
		return
	}
	user := &AuthUser{}
	user = GetAuthUser(username)
	if passwordModify.OldPassword != user.Password {
		c.JSON(http.StatusOK, MsgFail("当前密码有误"))
		return
	}
	user.Password = passwordModify.NewPassword
	SaveAuthUser(user)
	c.JSON(http.StatusOK, Ok())
}

func GetUserRolesByUserIdHandler(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	c.JSON(http.StatusOK, OkData(GetUserRolesByUserId(uint(userId))))
}

func AssignRolesHandler(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("userId"), 10, 32)
	roleIds := &[]uint{}
	c.ShouldBindJSON(roleIds)
	if err != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	SaveUserRoles(uint(userId), *roleIds)
	c.JSON(http.StatusOK, Ok())
}

func GetRolePageHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	keyword := c.Query("keyword")
	c.JSON(http.StatusOK, OkData(GetRolePage(pageParams.Page, pageParams.Size, keyword)))
}

func ListRoleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, OkData(ListAllRoles()))
}

func AddRoleHandler(c *gin.Context) {
	roleCreate := &RoleCreate{}
	c.ShouldBindJSON(roleCreate)
	if roleCreate.Key == "" || roleCreate.Name == "" {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	if GetRoleByKey(roleCreate.Key).ID != 0 {
		c.JSON(http.StatusOK, CodeFail(DataConflict))
	}
	role := &Role{}
	role.Key = roleCreate.Key
	role.Name = roleCreate.Name
	role.Description = roleCreate.Description
	SaveRole(role)
	c.JSON(http.StatusOK, Ok())
}

func UpdateRoleHandler(c *gin.Context) {
	roleUpdate := &RoleUpdate{}
	c.ShouldBindJSON(roleUpdate)
	if roleUpdate.ID == 0 || roleUpdate.Key == "" || roleUpdate.Name == "" {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	roleInDb := GetRoleByKey(roleUpdate.Key)
	if roleInDb.ID != 0 && roleInDb.ID != roleUpdate.ID {
		c.JSON(http.StatusOK, CodeFail(DataConflict))
		return
	}
	role := &Role{}
	role.ID = roleUpdate.ID
	role.Key = roleUpdate.Key
	role.Name = roleUpdate.Name
	role.Description = roleUpdate.Description
	SaveRole(role)
	c.JSON(http.StatusOK, Ok())
}

func DeleteRoleHandler(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("roleId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	DeleteRole(uint(roleId))
	c.JSON(http.StatusOK, Ok())
}

func GetRolePermissionsByRoleIdHandler(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("roleId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	c.JSON(http.StatusOK, OkData(GetRolePermissionsByRoleId(uint(roleId))))
}

func AssignPermissionsHandler(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Query("roleId"), 10, 32)
	permissionIds := &[]uint{}
	c.ShouldBindJSON(permissionIds)
	if err != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	SaveRolePermissions(uint(roleId), *permissionIds)
	c.JSON(http.StatusOK, Ok())
}

func GetPermissionPageHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	permissionQuery := &PermissionQuery{}
	c.ShouldBindJSON(permissionQuery)
	c.JSON(http.StatusOK, OkData(GetPermissionPage(pageParams.Page, pageParams.Size, permissionQuery)))
}

func ListPermissionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, OkData(ListAllPermissions()))
}

func AddPermissionHandler(c *gin.Context) {
	permissionCreate := &PermissionCreate{}
	c.ShouldBindJSON(permissionCreate)
	if permissionCreate.Key == "" || permissionCreate.Name == "" {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	if GetPermissionByKey(permissionCreate.Key).ID != 0 {
		c.JSON(http.StatusOK, CodeFail(DataConflict))
		return
	}
	permission := &Permission{}
	permission.Key = permissionCreate.Key
	permission.Name = permissionCreate.Name
	permission.PermType = permissionCreate.PermType
	permission.Description = permissionCreate.Description
	SavePermission(permission)
	c.JSON(http.StatusOK, Ok())
}

func UpdatePermissionHandler(c *gin.Context) {
	permissionUpdate := &PermissionUpdate{}
	c.ShouldBindJSON(permissionUpdate)
	if permissionUpdate.ID == 0 || permissionUpdate.Key == "" || permissionUpdate.Name == "" {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	permissionInDb := GetPermissionByKey(permissionUpdate.Key)
	if permissionInDb.ID != 0 && permissionInDb.ID != permissionUpdate.ID {
		c.JSON(http.StatusOK, CodeFail(DataConflict))
		return
	}
	permission := &Permission{}
	permission.ID = permissionUpdate.ID
	permission.Key = permissionUpdate.Key
	permission.Name = permissionUpdate.Name
	permission.PermType = permissionUpdate.PermType
	permission.Description = permissionUpdate.Description
	SavePermission(permission)
	c.JSON(http.StatusOK, Ok())
}

func DeletePermissionHandler(c *gin.Context) {
	permissionId, err := strconv.ParseUint(c.Param("permissionId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
		return
	}
	DeletePermission(uint(permissionId))
	c.JSON(http.StatusOK, Ok())
}
