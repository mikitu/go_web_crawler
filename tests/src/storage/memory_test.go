package tests
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/mikitu/go_web_crawler/src/storage"
)

func TestSetAndGet(t *testing.T) {

	mem_storage := storage.NewMemoryStorage()
	value := "123"
	Convey("Given value = " + value, t, func() {
		Convey("When call Set(\"value\", value)", func() {
			mem_storage.Set("value", value)
			Convey("Should receive '123' back from Get(\"value\")", func() {
				res := mem_storage.Get("value").(string)
				So(res, ShouldEqual, value )
			})
		})
	})
}

func TestGetAll(t *testing.T) {

	mem_storage := storage.NewMemoryStorage()
	mem_storage.Set("value", "123")
	mem_storage.Set("value_int", 123)
	expected := map[string]interface{}{"value": "123", "value_int" : 123}
	Convey("Saving multiple values", t, func() {
		Convey("Should return all values", func() {
			res := mem_storage.GetAll()
			So(res, ShouldHaveLength, 2)
			So(res, ShouldContainKey, "value")
			So(res, ShouldContainKey, "value_int")

			So(res["value"], ShouldEqual, expected["value"] )
			So(res["value_int"], ShouldEqual, expected["value_int"] )
		})
	})
}
func TestExists(t *testing.T) {

	mem_storage := storage.NewMemoryStorage()
	mem_storage.Set("key", "123")
	Convey("Should return \"true\"/\"false\" if the key \"key\" exists or not in storage", t, func() {
		So(mem_storage.Exists("key"), ShouldBeTrue )
		So(mem_storage.Exists("key_not_exists"), ShouldBeFalse )
	})
}

