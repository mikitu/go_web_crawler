package crawler

import "net/url"

type Options struct {
	data map[string]interface{}
}

func NewCrawlerOptions(_url string) *Options {
	opt := &Options{data: make(map[string]interface{})}
	opt.parseUrl(_url)
	return opt
}

func (o *Options) Set(key string, value interface{}) {
	o.data[key] = value
}
func (o Options) Get(key string) interface{} {
	if val, ok := o.data[key]; ok {
		return val
	}
	return nil
}

func (o *Options) parseUrl(base_url string) {
	_url, _ := url.Parse(base_url)

	if _url.Scheme != "" && _url.Host != "" {
		o.Set("scheme", _url.Scheme)
		o.Set("host", _url.Host)
		o.Set("base_url", _url.Scheme + "://" + _url.Host)
	} else {
		panic("Non standard url provided")
	}
}
