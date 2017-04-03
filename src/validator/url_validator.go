package validator

import (
	"strings"
	"net/url"
)

type UrlValidator struct {
	base_url string
	url *url.URL
}

func NewUrlValidator(base_url string) *UrlValidator {
	return &UrlValidator{base_url: base_url}
}

//validate url and return normalized url
func (v *UrlValidator) Validate(_url string) (valid bool, n_url string) {
	var err error
	v.url, err = url.Parse(_url)
	if err != nil {
		return false, ""
	}
	if v.isAbsolute() && ! v.sameHost() {
		return false, ""
	}

	if v.isAnchor() || ! v.isValidPath() {
		return false, ""
	}

	return true, v.normalize()
}

func (v UrlValidator) sameHost() bool  {
	u, _ := url.Parse(v.base_url)
	return u.Host == v.url.Host
}

func (v UrlValidator) isAnchor() bool  {
	return strings.Index(v.url.Path, "#") == 0
}

func (v UrlValidator) isEmail() bool  {
	return strings.Index(v.url.Path, "mailto:") == 0
}

func (v UrlValidator) isTel() bool  {
	return strings.Index(v.url.Path, "tel:") == 0
}

func (v UrlValidator) isValidPath() bool  {
	if v.isEmail() || v.isTel() {
		return false
	}
	if strings.Index(v.url.Path, "*") == 0 {
		return false
	}
	return true
}

func (v UrlValidator) isAbsolute() bool {
	return v.url.Scheme != "" && v.url.Host != ""
}

func (v UrlValidator) normalize() string {
	if v.isAbsolute() {
		return v.url.String()
	}
	startWithSlash := strings.Index(v.url.Path, "/") == 0
	if startWithSlash {
		return v.base_url + v.url.String()
	}
	return v.base_url + "/" + v.url.String()

}

