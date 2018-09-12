package yugle

import (
	"github.com/jinzhu/gorm"
)

type ShotOnOnePlusPicture struct {
	gorm.Model
	Picture string `gorm:"unique;not null"`
	Title   string `gorm:"not null"`
	Author  string `gorm:"not null"`
	Date    string `gorm:"unique;not null"`
}

type BingPicture struct {
	gorm.Model
	Picture string `gorm:"unique;not null"`
	Title   string `gorm:"not null"`
	Date    string `gorm:"unique;not null"`
}

func GetShotOnOnePlusPictures(page int, size int) *Pagination {
	db := DbConnect()
	defer db.Close()
	pictures := &[]ShotOnOnePlusPicture{}
	db.Limit(size).Offset((page - 1) * size).Order("date desc").Find(pictures)
	var total int
	db.Table("shot_on_one_plus_pictures").Count(&total)
	return GenPage(page, size, total, pictures)
}

func GetBingPictures(page int, size int) *Pagination {
	db := DbConnect()
	defer db.Close()
	pictures := &[]BingPicture{}
	db.Limit(size).Offset((page - 1) * size).Order("date desc").Find(pictures)
	var total int
	db.Table("bing_pictures").Count(&total)
	return GenPage(page, size, total, pictures)
}
