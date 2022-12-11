package pwd

import (
	"net/http"
	"time"
)

func NewParser(timeout time.Duration) Parser {
	p := &query{}
	p.setTimeOut(timeout)
	return p
}

type Parser interface {
	Parse(url string) (*http.Response, error)
	SetUserAgent(userAgent string)
	SetProxy(prx *Proxy)
}
