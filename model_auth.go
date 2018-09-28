package yugle

import (
	"github.com/jinzhu/gorm"
)

const (
	Action = iota
	Menu
	Button
)

type AuthUser struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null;default:123456"`
	Enabled  bool   `gorm:"default:true"`
}

type Role struct {
	gorm.Model
	Key         string `gorm:"unique;not null"`
	Name        string `gorm:"not null;index:idx_role_name"`
	Description string
}

type Permission struct {
	gorm.Model
	Key         string `gorm:"unique;not null"`
	Name        string `gorm:"not null;index:idx_perm_name"`
	PermType    int8   `gorm:"not null"`
	Description string
}

type UserRole struct {
	UserId uint `gorm:"not null"`
	RoleId uint `gorm:"not null"`
}

type RolePermission struct {
	RoleId       uint `gorm:"not null"`
	PermissionId uint `gorm:"not null"`
}

func GetAuthUserPage(page int, size int, query *AuthUserQuery) *Pagination {
	db := DbConnect()
	defer db.Close()
	users := &[]AuthUser{}
	var querySql string
	var total int
	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"
		querySql = "username LIKE ? AND enabled = ?"
		db.Model(users).Where(querySql, keyword, query.Enabled).Count(&total)
		db.Limit(size).Offset((page-1)*size).Where(querySql, keyword, query.Enabled).Order("created_at desc").Find(users)
	} else {
		querySql = "enabled = ?"
		db.Model(users).Where(querySql, query.Enabled).Count(&total)
		db.Limit(size).Offset((page-1)*size).Where(querySql, query.Enabled).Order("created_at desc").Find(users)
	}
	var userVOs []UserVO
	for _, user := range *users {
		userVO := UserVO{}
		userVO.Username = user.Username
		userVO.Enabled = user.Enabled
		userVO.User = *GetUserById(user.ID)
		userVO.User.UserId = user.ID
		userVOs = append(userVOs, userVO)
	}
	return GenPage(page, size, total, userVOs)
}

func GetAuthUser(username string) *AuthUser {
	db := DbConnect()
	defer db.Close()
	user := AuthUser{}
	db.Where("username = ?", username).Find(&user)
	return &user
}

func GetAuthUserById(userId uint) *AuthUser {
	db := DbConnect()
	defer db.Close()
	user := AuthUser{}
	db.Where("id = ?", userId).Find(&user)
	return &user
}

func SaveAuthUser(user *AuthUser) {
	db := DbConnect()
	defer db.Close()
	db.Save(user)
}

func GetUserRolesByUserId(userId uint) *[]uint {
	db := DbConnect()
	defer db.Close()
	userRoles := &[]UserRole{}
	db.Where("user_id = ?", userId).Find(userRoles)
	roleIds := &[]uint{}
	for _, userRole := range *userRoles {
		*roleIds = append(*roleIds, userRole.RoleId)
	}
	return roleIds
}

func SaveUserRoles(userId uint, roleIds []uint) {
	db := DbConnect()
	defer db.Close()
	db.Where("user_id = ?", userId).Delete(UserRole{})
	for _, roleId := range roleIds {
		userRole := &UserRole{userId, roleId}
		db.Save(userRole)
	}
}

func GetRolePage(page int, size int, keyword string) *Pagination {
	db := DbConnect()
	defer db.Close()
	roles := &[]Role{}
	var querySql string
	var total int
	if keyword != "" {
		keyword = "%" + keyword + "%"
		querySql = "name LIKE ? OR description LIKE ?"
		db.Model(roles).Where(querySql, keyword, keyword).Count(&total)
		db.Limit(size).Offset((page-1)*size).Where(querySql, keyword, keyword).Order("created_at desc").Find(roles)
	} else {
		db.Model(roles).Count(&total)
		db.Limit(size).Offset((page - 1) * size).Order("created_at desc").Find(roles)
	}
	return GenPage(page, size, total, roles)
}

func ListAllRoles() *[]RoleVO {
	db := DbConnect()
	defer db.Close()
	roles := &[]Role{}
	db.Model(roles).Find(roles)
	roleVOs := &[]RoleVO{}
	for _, role := range *roles {
		roleVO := RoleVO{}
		roleVO.ID = role.ID
		roleVO.Key = role.Key
		roleVO.Name = role.Name
		roleVO.Description = role.Description
		roleVO.CreatedAt = role.CreatedAt
		*roleVOs = append(*roleVOs, roleVO)
	}
	return roleVOs
}

func GetRoleByKey(key string) *Role {
	db := DbConnect()
	defer db.Close()
	role := Role{}
	db.Where("key = ?", key).Find(&role)
	return &role
}

func SaveRole(role *Role) {
	db := DbConnect()
	defer db.Close()
	db.Save(role)
}

func GetRolePermissionsByRoleId(roleId uint) *[]uint {
	db := DbConnect()
	defer db.Close()
	rolePermissions := &[]RolePermission{}
	db.Where("role_id = ?", roleId).Find(rolePermissions)
	permissionIds := &[]uint{}
	for _, rolePermission := range *rolePermissions {
		*permissionIds = append(*permissionIds, rolePermission.PermissionId)
	}
	return permissionIds
}

func SaveRolePermissions(roleId uint, permissionIds []uint) {
	db := DbConnect()
	defer db.Close()
	db.Where("role_id = ?", roleId).Delete(RolePermission{})
	for _, permissionId := range permissionIds {
		rolePermission := &RolePermission{roleId, permissionId}
		db.Save(rolePermission)
	}
}

func ListAllPermissions() *[]PermissionVO {
	db := DbConnect()
	defer db.Close()
	permissions := &[]Permission{}
	db.Model(permissions).Find(permissions)
	permissionVOs := &[]PermissionVO{}
	for _, permission := range *permissions {
		permissionVO := PermissionVO{}
		permissionVO.ID = permission.ID
		permissionVO.Key = permission.Key
		permissionVO.Name = permission.Name
		permissionVO.PermType = permission.PermType
		permissionVO.Description = permission.Description
		permissionVO.CreatedAt = permission.CreatedAt
		*permissionVOs = append(*permissionVOs, permissionVO)
	}
	return permissionVOs
}
