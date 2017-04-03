package parser

import (
	"bytes"
	"github.com/mikitu/go_web_crawler/src/validator"
	"regexp"
)

type RegexParser struct {
	validator validator.ValidatorInterface
}

func NewRegexParser(_validator validator.ValidatorInterface) *RegexParser {
	return &RegexParser{validator: _validator}
}

func (p RegexParser) Parse(body *bytes.Buffer, ch chan interface{}, stop chan bool) {
	var re = regexp.MustCompile(`<a[^>]+href="([^"]+)"[^>]+?>`)

	for _, match := range re.FindAllStringSubmatch(body.String(), -1) {
		is_valid, _url := p.validator.Validate(match[1])
		if  is_valid {
			ch <- _url
		}
	}
	stop <- true
}
