package tests

import (
	"testing"
	"github.com/mikitu/go_web_crawler/src/validator"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateUrl(t *testing.T) {
	base_url := "http://golangweekly.com"
	v := validator.NewUrlValidator(base_url)

	Convey("Given an url", t, func() {
		Convey("When url is invalid", func() {
			_url := "invalid://Url"
			is_valid, _ := v.Validate(_url)
			So(is_valid, ShouldBeFalse)
		})
		Convey("When Url is valid", func() {
			_url := "http://golangweekly.com/123"
			is_valid, _ := v.Validate(_url)
			So(is_valid, ShouldBeTrue)
		})
		Convey("When url is not in the same domain", func() {
			_url := "http://linux.com/news/123"
			is_valid, _ := v.Validate(_url)
			So(is_valid, ShouldBeFalse)
		})
		Convey("When url is absolute and start with /", func() {
			_url := "/news/123"
			is_valid, n_url := v.Validate(_url)
			So(is_valid, ShouldBeTrue)
			So(n_url, ShouldEqual, base_url + _url)
		})
		Convey("When url is absolute and doesn't start with /", func() {
			_url := "news/123"
			is_valid, n_url := v.Validate(_url)
			So(is_valid, ShouldBeTrue)
			So(n_url, ShouldEqual, base_url + "/" + _url)
		})
		Convey("When url is email", func() {
			_url := "mailto:test@golang.org"
			is_valid, _ := v.Validate(_url)
			So(is_valid, ShouldBeFalse)
		})
	})

}
