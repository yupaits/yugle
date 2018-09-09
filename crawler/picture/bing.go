package crawler

import (
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/page_processer"
	"github.com/hu17889/go_spider/core/spider"
	"log"
)

type BingPictureProcessor struct {
	processor page_processer.PageProcesser
}

func NewBingPictureProcessor() *BingPictureProcessor {
	return &BingPictureProcessor{}
}

func (spider *BingPictureProcessor) Process(p *page.Page) {
	log.Println(p)
}

func (spider *BingPictureProcessor) Finish() {
	log.Println("Finish")
}

func CrawlBingPicture() {
	spider.NewSpider(NewBingPictureProcessor(), "bing_picture").AddUrl("https://cn.bing.com/", "html").Run()
}
