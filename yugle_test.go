package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	"yugle/crawler/picture"
	"yugle/dbutils"
	"yugle/model"
)

func TestAddAuthUser(t *testing.T) {
	db := dbutils.Connect()
	defer db.Close()
	authUser1 := &model.AuthUser{Username: "admin", Password: "123456"}
	authUser2 := &model.AuthUser{Model: gorm.Model{ID: 4}, Username: "user", Password: "123456"}
	db.DropTableIfExists(authUser1)
	db.CreateTable(authUser1)
	if db.NewRecord(authUser1) {
		db.Create(authUser1)
	}
	if db.NewRecord(authUser2) {
		db.Create(authUser2)
	}
}

func TestCrawlBingPicture(t *testing.T) {
	//db := dbutils.Connect()
	//defer db.Close()
	//picture := &crawler.BingPicture{}
	//db.DropTableIfExists(picture)
	//db.CreateTable(picture)
	crawler.CrawlBingPicture()
	log.Println(crawler.GetBingPictures(1, 3))
}

func TestCrawlShotOnOnePlus(t *testing.T) {
	db := dbutils.Connect()
	defer db.Close()
	picture := &crawler.ShotOnOnePlusPicture{}
	db.CreateTable(picture)
	crawler.CrawlShotOnOnePlusPicture()
}
