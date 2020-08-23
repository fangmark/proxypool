package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/zu1k/proxypool/internal/app"
)

func Cron() {
	_ = gocron.Every(10).Minutes().Do(app.CrawlGo)
	<-gocron.Start()
}
