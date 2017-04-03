package storage

import (
	"github.com/mikitu/go_web_crawler/src/parser"
	"bytes"
	"github.com/mikitu/go_web_crawler/src/validator"
)

type ImagesStorage struct {
	parser *parser.RegexImagesParser
	data map[string][]string
}

func NewImagesStorage() *ImagesStorage {
	return &ImagesStorage{
		parser.NewRegexImagesParser(validator.NewImageValidator()),
		map[string][]string{"urls": {}, "images": {}},
	}
}

func (s *ImagesStorage) Set(key string, data interface{}) {
	s.data["urls"] = append(s.data["urls"], key)
	s.getImages(data.(*bytes.Buffer))
}

func (s ImagesStorage) Get(key string) interface{} {
	if value, ok := s.data[key]; ok {
		return value
	}
	return nil
}
func (s ImagesStorage) GetAll()(map[string]interface{}) {
	ret := make(map[string]interface{})
	ret["urls"] = s.data["urls"]
	ret["images"] = s.data["images"]
	return ret
}

func (s ImagesStorage) Exists(key string) (bool) {
	return s.existsIn(key, "urls")
}

func (s ImagesStorage) imageExists(key string) (bool) {
	return s.existsIn(key, "images")
}

func (s ImagesStorage) existsIn(key, master string) (bool) {
	for _, _url := range s.data[master] {
		if key == _url {
			return true
		}
	}

	return false
}

func (s *ImagesStorage) getImages(body *bytes.Buffer) {
	ch := make(chan interface{})
	stop := make(chan bool)
	go s.parser.Parse(body, ch, stop)
	for {
		select  {
		case _url := <-ch:
			if ! s.imageExists(_url.(string)) {
				s.addImage(_url.(string))
			}
		case <-stop:
			return
		}
	}
}

func (s *ImagesStorage) addImage(_url string) {
	s.data["images"] = append(s.data["images"], _url)
}