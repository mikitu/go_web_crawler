package parser

import (
	"bytes"
	"regexp"
	"github.com/mikitu/go_web_crawler/src/validator"
)

type RegexImagesParser struct {
	validaor validator.ValidatorInterface
}

func NewRegexImagesParser(validaor validator.ValidatorInterface) *RegexImagesParser {
	return &RegexImagesParser{validaor: validaor}
}

func (p RegexImagesParser) Parse(body *bytes.Buffer, ch chan interface{}, stop chan bool) {
	var re_images = regexp.MustCompile(`<img[^>]+src="([^"]+)"[^>]+?>`)
	for _, match := range re_images.FindAllStringSubmatch(body.String(), -1) {
		is_valid, _url := p.validaor.Validate(match[1])
		if is_valid {
			ch <- _url
		}
	}
	stop <- true
}
