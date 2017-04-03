package fetcher

import (
	"net/http"
	"bytes"
)

type DefaultHttpFetcher struct {}
func NewDefaultHttpFetcher() * DefaultHttpFetcher {
	return &DefaultHttpFetcher{}
}

func (f DefaultHttpFetcher) Fetch(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	return buf, nil
}