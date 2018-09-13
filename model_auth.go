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
	Password string `gorm:"not null"`
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
		querySql = "username Like ? AND enabled = ?"
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

func SaveAuthUser(user *AuthUser) {
	db := DbConnect()
	defer db.Close()
	if db.NewRecord(user) {
		db.Create(user)
	}
}
