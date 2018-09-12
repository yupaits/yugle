package yugle

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/page_processer"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
	"time"
)

var bingPictureUrl = "https://cn.bing.com"
var BingPictureTaskName = "Bing壁纸爬虫"

type BingPictureProcessor struct {
	processor page_processer.PageProcesser
}

type BingPicturePipeline struct {
}

func NewBingPictureProcessor() *BingPictureProcessor {
	return &BingPictureProcessor{}
}

func (spider *BingPictureProcessor) Process(p *page.Page) {
	//log.Println(p)
	query := p.GetHtmlParser()
	//获取必应每日壁纸链接
	query.Find(`img[style="display:none"]`).Each(func(i int, selection *goquery.Selection) {
		img, imgExists := selection.Attr("src")
		if imgExists && strings.Index(img, "/az/hprichbg/rb") == 0 {
			p.AddField("picture", bingPictureUrl+img)
		}
	})
	//获取壁纸标题
	title, titleExists := query.Find(`a[id="sh_cp"]`).First().Attr("title")
	if titleExists {
		p.AddField("title", title)
	}
}

func (spider *BingPictureProcessor) Finish() {
	//用于接口实现声明，无实际业务
}

func NewBingPicturePipeline() *BingPicturePipeline {
	return &BingPicturePipeline{}
}

func (pipeline *BingPicturePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	db := DbConnect()
	defer db.Close()
	picture, pictureOk := items.GetItem("picture")
	title, titleOk := items.GetItem("title")
	if pictureOk && titleOk {
		bingPicture := &BingPicture{Picture: picture, Title: title, Date: time.Now().Local().String()[:10]}
		//查询当前爬取的图片是否已经爬取过，如果已经爬取过则不会重复保存
		bingPictureInDb := &BingPicture{}
		db.Where("picture = ?", picture).Find(bingPictureInDb)
		if bingPictureInDb.Picture == "" {
			db.Create(bingPicture)
		}
	}
}

func CrawlBingPicture() {
	bingTask := GetTaskByTaskName(BingPictureTaskName)
	bingTask.State = RUNNING
	SaveTask(bingTask)
	spider.NewSpider(NewBingPictureProcessor(), "bing_picture").
		AddUrl(bingPictureUrl, "html").
		AddPipeline(NewBingPicturePipeline()).
		SetThreadnum(1).
		Run()
}
