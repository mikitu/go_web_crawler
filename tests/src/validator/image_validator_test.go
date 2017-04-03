package tests

import (
	"testing"
	"github.com/mikitu/go_web_crawler/src/validator"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateImage(t *testing.T) {
	v := validator.NewImageValidator()

	// Only pass t into top-level Convey calls
	Convey("Given an image url or filename", t, func() {
		Convey("When validate jpg", func() {
			image := "sample.jpg"
			is_valid, _ := v.Validate(image)
			So(is_valid, ShouldBeTrue)
		})
		Convey("When validate png", func() {
			image := "sample.png"
			is_valid, _ := v.Validate(image)
			So(is_valid, ShouldBeTrue)
		})
		Convey("When validate gif", func() {
			image := "sample.gif"
			is_valid, _ := v.Validate(image)
			So(is_valid, ShouldBeTrue)
		})
		Convey("Otherwise", func() {
			image := "sample.giff"
			is_valid, _ := v.Validate(image)
			So(is_valid, ShouldBeFalse)
		})
	})

}
