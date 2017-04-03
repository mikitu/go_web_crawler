package validator

import (
	"net/url"
	"regexp"
)

type ImageValidator struct {
	base_url string
	url *url.URL
}

func NewImageValidator() *ImageValidator {
	return &ImageValidator{}
}

//validate image
func (v *ImageValidator) Validate(_url string) (valid bool, n_url string) {
	var re = regexp.MustCompile(`\.(jpg|gif|png)$`)
	if re.FindString(_url) == "" {
		return false, _url
	}
	return true, _url
}

