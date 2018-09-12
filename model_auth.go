package yugle

import (
	"github.com/jinzhu/gorm"
)

type AuthUser struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Enabled  bool   `gorm:"default:true"`
}

type Role struct {
}

type Permission struct {
}

func GetAuthUser(username string) *AuthUser {
	db := DbConnect()
	defer db.Close()
	user := AuthUser{}
	db.Where("username = ?", username).Find(&user)
	return &user
}
