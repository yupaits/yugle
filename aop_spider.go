package yugle

import (
	"github.com/gogap/aop"
	"github.com/hu17889/go_spider/core/spider"
)

type AppAop struct {
	spiderAop aop.AOP
}

func (this *AppAop) InitAop() {
	factory := aop.NewClassicBeanFactory()
	factory.RegisterBean("spider", new(spider.Spider))

	aspect := aop.NewAspect("close_callback", "spider")
	aspect.SetBeanFactory(factory)

	pointcut := aop.NewPointcut("spider_close_pointcut").Execution(`close()`)

	aspect.AddPointcut(pointcut)

	aspect.AddAdvice(&aop.Advice{Ordering: aop.After, Method: "SpiderCompleteCallback", PointcutRefID: "spider_close_pointcut"})

	this.spiderAop = *aop.NewAOP()
	this.spiderAop.SetBeanFactory(factory)
	this.spiderAop.AddAspect(aspect)
}

func RegisterAop() {
	appAop := AppAop{}
	appAop.InitAop()
}

func SpiderCompleteCallback() {

}
