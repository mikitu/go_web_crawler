package main

import (
	"github.com/mikitu/go_web_crawler/src/crawler"
	"fmt"
	"github.com/mikitu/go_web_crawler/src/parser"
	"github.com/mikitu/go_web_crawler/src/validator"
	log "github.com/sirupsen/logrus"

)
func init() {
	log.SetLevel(log.DebugLevel)
}
func main()  {
	//_url := "http://dermatocosmetice.eu"
	_url := "http://golangweekly.com"
	//_url := "https://www.linux.com"
	options := crawler.NewCrawlerOptions(_url)

	url_validator := validator.NewUrlValidator(_url)
	_parser := parser.NewRegexParser(url_validator)
	options.Set("parser", _parser)
	cr := crawler.NewCrawler(options)
	cr.Run(_url, 5)
	for u, _ := range cr.GetStorage().GetAll() {
		fmt.Printf("%+v\n", u)
	}
}