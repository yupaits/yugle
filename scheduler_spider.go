package yugle

import (
	"github.com/robfig/cron"
	"log"
)

const (
	bingPictureSpec   = "0 45 9 * * *"  //每天早上六点
	shotOnOnePlusSpec = "0 23 17 * * *" //每天早上六点
)

func StartSpiderScheduler() {
	//必应壁纸
	bingCron := cron.New()
	bingCron.AddFunc(bingPictureSpec, CrawlBingPicture)
	bingCron.Start()
	bingTask := InsertTaskIfAbsent(BingPictureTaskName)
	bingEntry := bingCron.Entries()[0]
	if bingEntry != nil {
		bingTask.Prev = bingEntry.Prev
		bingTask.Next = bingEntry.Next
		SaveTask(bingTask)
	}
	log.Println("Bing picture spider started.")

	//Shot on OnePlus壁纸
	shotOnOnePlusCron := cron.New()
	shotOnOnePlusCron.AddFunc(shotOnOnePlusSpec, CrawlShotOnOnePlusPicture)
	shotOnOnePlusCron.Start()
	shotOnOnePlusTask := InsertTaskIfAbsent(ShotOnOnePlusTaskName)
	shotOnOnePlusEntry := shotOnOnePlusCron.Entries()[0]
	if shotOnOnePlusEntry != nil {
		shotOnOnePlusTask.Prev = shotOnOnePlusEntry.Prev
		shotOnOnePlusTask.Next = shotOnOnePlusEntry.Next
		SaveTask(shotOnOnePlusTask)
	}
	log.Println("Shot on OnePlus spider started.")
}
