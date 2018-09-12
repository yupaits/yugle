package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	"yugle/crawler/picture"
	"yugle/dbutils"
	"yugle/model"
	"yugle/scheduler/task"
)

func TestCreateTables(t *testing.T) {
	db := dbutils.Connect()
	defer db.Close()

	//authUser := &model.AuthUser{}
	//db.DropTableIfExists(authUser)
	//db.CreateTable(authUser)
	//
	//bingPicture := &crawler.BingPicture{}
	//db.DropTableIfExists(bingPicture)
	//db.CreateTable(bingPicture)
	//
	//onePlusPicture := &crawler.ShotOnOnePlusPicture{}
	//db.DropTableIfExists(onePlusPicture)
	//db.CreateTable(onePlusPicture)

	taskExample := &task.Task{}
	db.DropTableIfExists(taskExample)
	db.CreateTable(taskExample)
}

func TestAddAuthUser(t *testing.T) {
	db := dbutils.Connect()
	defer db.Close()
	authUser1 := &model.AuthUser{Username: "admin", Password: "123456"}
	authUser2 := &model.AuthUser{Model: gorm.Model{ID: 4}, Username: "user", Password: "123456"}
	if db.NewRecord(authUser1) {
		db.Create(authUser1)
	}
	if db.NewRecord(authUser2) {
		db.Create(authUser2)
	}
}

func TestCrawlBingPicture(t *testing.T) {
	crawler.CrawlBingPicture()
	log.Println(crawler.GetBingPictures(1, 3))
}

func TestCrawlShotOnOnePlus(t *testing.T) {
	crawler.CrawlShotOnOnePlusPicture()
}
