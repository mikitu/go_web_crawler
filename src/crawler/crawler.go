package crawler

import (
	"fmt"
	"github.com/mikitu/go_web_crawler/src/parser"
	"github.com/mikitu/go_web_crawler/src/storage"
	"github.com/mikitu/go_web_crawler/src/fetcher"
	"github.com/mikitu/go_web_crawler/src/validator"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type Crawler struct {
	ch chan interface{}
	stop chan bool
	fetcher fetcher.Fetcher
	parser parser.Parser
	storage storage.Storage
	options *Options
}

func NewCrawler(options *Options) *Crawler{

	cr := new(Crawler)
	cr.options = options
	cr.init()
	return cr
}

// initialize default options for crawler
func (c *Crawler) init () {
	_parser := c.getOption("parser")
	if _parser == nil {
		c.setupDefaultParser()
	} else {
		c.parser = _parser.(parser.Parser)
	}
	_fetcher := c.getOption("fetcher")
	if _fetcher == nil {
		c.fetcher = fetcher.NewDefaultHttpFetcher()
	} else {
		c.fetcher = _fetcher.(fetcher.Fetcher)
	}
	_storage := c.getOption("storage")
	if _storage == nil {
		c.storage = storage.NewMemoryStorage()
	} else {
		c.storage = _storage.(storage.Storage)
	}
	c.ch = make(chan interface{}, 10)
	c.stop = make(chan bool)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func(){
		<-ch
		log.Warn("Kil signal has been received")
		os.Exit(1)
	}()
}

func (c *Crawler) setupDefaultParser() {
	_validator := validator.NewUrlValidator(c.options.Get("base_url").(string))
	c.parser = parser.NewDefaultParser(_validator)
}

func (c Crawler) GetStorage() storage.Storage {
	return c.storage
}

func (c Crawler) getOption (key string) interface{} {
	return c.options.Get(key)
}

func (c *Crawler) setOption (key string, value interface{}) {
	c.options.Set(key, value)
}


func (c Crawler) visited (url string) bool {
	return c.storage.Exists(url)
}

//recursively crawl pages starting with url, to a maximum of depth.

func (c Crawler) Run (url string, depth int) {
	if depth <= 0 {
		return
	}
	body, err := c.fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	go c.parser.Parse(body, c.ch, c.stop)
	c.storage.Set(url, body)
	for {
		select {
		case _url := <-c.ch:
			if ! c.visited(_url.(string)) {
				log.Debug(_url)
				c.Run(_url.(string), depth-1)
			}
		case <-c.stop:
			return
		}
	}
	return
}
