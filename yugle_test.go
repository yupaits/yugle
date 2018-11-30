package yugle

import (
	"log"
	"testing"
)

func TestCreateTables(t *testing.T) {
	db := DbConnect()
	defer db.Close()

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

func TestCrawlBingPicture(t *testing.T) {
	CrawlBingPicture()
	log.Println(GetBingPictures(1, 3))
}

func TestCrawlShotOnOnePlus(t *testing.T) {
	CrawlShotOnOnePlusPicture()
}
