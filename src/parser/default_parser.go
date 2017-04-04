package parser

import (
	"golang.org/x/net/html"
	"bytes"
	"github.com/mikitu/go_web_crawler/src/validator"
)

type DefaultParser struct {
	validator validator.ValidatorInterface
}

func NewDefaultParser(_validator validator.ValidatorInterface) *DefaultParser {
	return &DefaultParser{validator: _validator}
}

func (p DefaultParser) Parse(body *bytes.Buffer, ch chan interface{}, stop chan bool) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			stop <- true
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, _url := p.getHref(t)
			if !ok {
				continue
			}
			is_valid, _url := p.validator.Validate(_url)
			if  is_valid {
				ch <- _url
			}
		}
	}
}

func (p DefaultParser) getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
			return
		}
	}
	return
}