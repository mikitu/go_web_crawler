golang web crawler
===========

A concurrent web crawler written in Go

#### Features

 -  Concurrent execution using goroutines
 -  Open, customisable design providing hooks into the execution logic
 -  Options for different storage

#### Installation and dependencies
Dependencies are managed with [glide](https://github.com/Masterminds/glide)

    go get -v github.com/Masterminds/glide
    glide install

#### Examples (./src/examples)

Default example (src/main.go) is using crawler defaults:

    _storage := storage.NewMemoryStorage()
    _validator := NewUrlValidator(url)
    _body_parser := parser.NewDefaultParser(_validator)
    _http_fetcher := fetcher.NewDefaultHttpFetcher()



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

##### Customisation
1. Providing a custom body parser. some examples are provided:
	- src/parser/default_parser.go
	- src/parser/regex_parser.go
	- src/parser/regex_images_parser.go

Crawler accepts any custom parser that implements parser.ParserInterface
file: src/examples/simple_regexp/main.go

	    func main()  {
			_url := "http://golangweekly.com"
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

2. Providing a custom storage option. some examples are provided:
	- src/parser/memory.go
	- src/parser/images_storage.go

Crawler accepts any custom storage that implements storage.Storage Interface
file: src/examples/url_and_images/main.go

	    func main()  {
			_url := "http://golangweekly.com"
			options := crawler.NewCrawlerOptions(_url)

			url_validator := validator.NewUrlValidator(_url)
			_parser := parser.NewRegexParser(url_validator)
			options.Set("parser", _parser)

			_storage := storage.NewImagesStorage()
			options.Set("storage", _storage)

			cr := crawler.NewCrawler(options)
			cr.Run(_url, 5)
			for u, _ := range cr.GetStorage().GetAll() {
				fmt.Printf("%+v\n", u)
			}
		}

3. Providing a custom way to fetch any url. any object that implements fetcher.Fetcher Interface is accepted
 example:
	- src/fetcher/default_http_fetcher.go

