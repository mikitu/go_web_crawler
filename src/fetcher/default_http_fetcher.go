package fetcher

import (
	"net/http"
	"bytes"
	log "github.com/sirupsen/logrus"
	"errors"
)

type DefaultHttpFetcher struct {}
func NewDefaultHttpFetcher() * DefaultHttpFetcher {
	return &DefaultHttpFetcher{}
}

func (f DefaultHttpFetcher) Fetch(url string) (*bytes.Buffer, error) {
	// TODO: provide a way to delay the requests after a number
	// of consecutive requests. some websites responds with
	// status 429 Too many requests
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Error("Something is wrong with the remote server")
		log.Errorf("StatusCode:  %d\n", resp.StatusCode )
		return nil, errors.New("Error: StatusCode is  " + resp.Status)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	return buf, nil
}