package fetcher

import "bytes"

type Fetcher interface {
	// Fetch returns the body of URL or error
	Fetch(url string) (*bytes.Buffer, error)
}
