package main

import (
	"github.com/mikitu/go_web_crawler/src/crawler"
	"flag"
	log "github.com/sirupsen/logrus"

	"os"
)
func init() {
	log.SetLevel(log.DebugLevel)
}
func main() {

	var url = flag.String("url", "", "url to crawl")
	flag.Parse()

	if *url == "" {
		log.Error("Please provide a valid url as -url parameter. Eg: -url=https://golangweekly.com/")
		os.Exit(0)
	}

	options := crawler.NewCrawlerOptions(*url)

	cr := crawler.NewCrawler(options)
	cr.Run(*url, 5)
	for url := range cr.GetStorage().GetAll() {
		log.Info(url)
	}
}

