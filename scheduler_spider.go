package yugle

import (
	"log"

	"github.com/robfig/cron"
)

const (
	bingPictureSpec   = "0 20 16 * * *" //每天早上六点
	shotOnOnePlusSpec = "0 20 16 * * *" //每天早上六点
)

var crons = make(map[string]*cron.Cron)

func StartSpiderScheduler() {
	//必应壁纸
	bingCron := cron.New()
	bingCron.AddFunc(bingPictureSpec, CrawlBingPicture)
	bingCron.Start()
	bingTask := InsertTaskIfAbsent(BingPictureTaskName)
	UpdateTask(bingTask, bingCron)
	crons[BingPictureTaskName] = bingCron
	log.Println("Bing picture spider started.")

	//Shot on OnePlus壁纸
	shotOnOnePlusCron := cron.New()
	shotOnOnePlusCron.AddFunc(shotOnOnePlusSpec, CrawlShotOnOnePlusPicture)
	shotOnOnePlusCron.Start()
	shotOnOnePlusTask := InsertTaskIfAbsent(ShotOnOnePlusTaskName)
	UpdateTask(shotOnOnePlusTask, shotOnOnePlusCron)
	crons[ShotOnOnePlusTaskName] = shotOnOnePlusCron
	log.Println("Shot on OnePlus spider started.")
}
