package scheduler

import (
	"github.com/robfig/cron"
	"log"
	"yugle/crawler/picture"
	"yugle/scheduler/task"
)

const (
	bingPictureSpec   = "0 45 9 * * *"  //每天早上六点
	shotOnOnePlusSpec = "0 23 17 * * *" //每天早上六点
)

func StartSpiderScheduler() {
	//必应壁纸
	bingCron := cron.New()
	bingCron.AddFunc(bingPictureSpec, crawler.CrawlBingPicture)
	bingCron.Start()
	bingTask := task.InsertTaskIfAbsent(crawler.BingPictureTaskName)
	bingEntry := bingCron.Entries()[0]
	if bingEntry != nil {
		bingTask.Prev = bingEntry.Prev
		bingTask.Next = bingEntry.Next
		task.SaveTask(bingTask)
	}
	log.Println("Bing picture spider started.")

	//Shot on OnePlus壁纸
	shotOnOnePlusCron := cron.New()
	shotOnOnePlusCron.AddFunc(shotOnOnePlusSpec, crawler.CrawlShotOnOnePlusPicture)
	shotOnOnePlusCron.Start()
	shotOnOnePlusTask := task.InsertTaskIfAbsent(crawler.ShotOnOnePlusTaskName)
	shotOnOnePlusEntry := shotOnOnePlusCron.Entries()[0]
	if shotOnOnePlusEntry != nil {
		shotOnOnePlusTask.Prev = shotOnOnePlusEntry.Prev
		shotOnOnePlusTask.Next = shotOnOnePlusEntry.Next
		task.SaveTask(shotOnOnePlusTask)
	}
	log.Println("Shot on OnePlus spider started.")
}
