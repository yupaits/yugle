package scheduler

import (
	"github.com/robfig/cron"
	"log"
	"yugle/crawler/picture"
)

const (
	bingPictureSpec = "0 55 15 * * *" //每天早上六点
)

func StartSpiderScheduler() {
	//必应壁纸任务
	bingCron := cron.New()
	bingCron.AddFunc(bingPictureSpec, crawler.CrawlBingPicture)
	bingCron.Start()
	log.Println("Bing picture spider started.")

}
