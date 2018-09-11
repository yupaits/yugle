package scheduler

import (
	"github.com/robfig/cron"
	"log"
	"yugle/crawler/picture"
)

const (
	bingPictureSpec   = "0 0 0 * * *"    //每天早上六点
	ShotOnOnePlusSpec = "10 19 14 * * *" //每天早上六点
)

func StartSpiderScheduler() {
	//必应壁纸
	bingCron := cron.New()
	bingCron.AddFunc(bingPictureSpec, crawler.CrawlBingPicture)
	bingCron.Start()
	log.Println("Bing picture spider started.")

	//Shot on OnePlus壁纸
	shotOnOnePlusCron := cron.New()
	shotOnOnePlusCron.AddFunc(ShotOnOnePlusSpec, crawler.CrawlShotOnOnePlusPicture)
	shotOnOnePlusCron.Start()
	log.Println("Shot on OnePlus spider started.")
}
