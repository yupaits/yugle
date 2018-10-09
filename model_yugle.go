package yugle

import "github.com/jinzhu/gorm"

const (
	Secret = iota
	Male
	Female
)

type User struct {
	gorm.Model
	UserId     uint
	Phone      string
	Email      string
	Gender     int8
	Avatar     string
	BirthYear  int
	BirthMonth int8
	BirthDay   int8
}

func GetUserById(userId uint) *User {
	db := DbConnect()
	defer db.Close()
	user := &User{}
	db.Where("user_id = ?", userId).Find(user)
	return user
}

func SaveUser(user *User) {
	db := DbConnect()
	defer db.Close()
	if db.NewRecord(user) {
		db.Create(user)
	} else {
		db.Model(user).Updates(user)
	}
}
