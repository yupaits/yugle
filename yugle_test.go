package yugle

import (
	"github.com/jinzhu/gorm"
	"log"
	"testing"
)

func TestCreateTables(t *testing.T) {
	db := DbConnect()
	defer db.Close()

	//authUser := &AuthUser{}
	//db.DropTableIfExists(authUser)
	//db.CreateTable(authUser)
	//db.AutoMigrate(authUser)

	//role := &Role{}
	//db.DropTableIfExists(role)
	//db.CreateTable(role)
	//db.AutoMigrate(role)

	//permission := &Permission{}
	//db.DropTableIfExists(permission)
	//db.CreateTable(permission)
	//db.AutoMigrate(permission)

	//userRole := &UserRole{}
	//db.DropTableIfExists(userRole)
	//db.CreateTable(userRole)

	//rolePermission := &RolePermission{}
	//db.DropTableIfExists(rolePermission)
	//db.CreateTable(rolePermission)

	//user := &User{}
	//db.DropTableIfExists(user)
	//db.CreateTable(user)

	//bingPicture := &BingPicture{}
	//db.DropTableIfExists(bingPicture)
	//db.CreateTable(bingPicture)
	//db.AutoMigrate(bingPicture)

	//onePlusPicture := &ShotOnOnePlusPicture{}
	//db.DropTableIfExists(onePlusPicture)
	//db.CreateTable(onePlusPicture)
	//db.AutoMigrate(onePlusPicture)

	//taskExample := &Task{}
	//db.DropTableIfExists(taskExample)
	//db.CreateTable(taskExample)
}

func TestAddAuthUser(t *testing.T) {
	db := DbConnect()
	defer db.Close()
	authUser1 := &AuthUser{Username: "admin", Password: "123456"}
	authUser2 := &AuthUser{Model: gorm.Model{ID: 4}, Username: "user", Password: "123456"}
	if db.NewRecord(authUser1) {
		db.Create(authUser1)
	}
	if db.NewRecord(authUser2) {
		db.Create(authUser2)
	}
}

func TestCrawlBingPicture(t *testing.T) {
	CrawlBingPicture()
	log.Println(GetBingPictures(1, 3))
}

func TestCrawlShotOnOnePlus(t *testing.T) {
	CrawlShotOnOnePlusPicture()
}
