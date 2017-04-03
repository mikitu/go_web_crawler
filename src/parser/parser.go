package parser

import (
	"bytes"
)

type Parser interface {
	Parse(body *bytes.Buffer, ch chan interface{}, stop chan bool)
}
