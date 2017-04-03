package main

import (
	"github.com/mikitu/go_web_crawler/src/crawler"
	"flag"
	log "github.com/sirupsen/logrus"

)
func init() {
	log.SetLevel(log.DebugLevel)
}
func main() {

	var url = flag.String("url", "", "url to crawl")
	flag.Parse()

	options := crawler.NewCrawlerOptions(*url)

	cr := crawler.NewCrawler(options)
	cr.Run(*url, 5)
	for url, _ := range cr.GetStorage().GetAll() {
		log.Info(url)
	}
}

