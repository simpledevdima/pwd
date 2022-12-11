package pwd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type query struct {
	timeout   time.Duration
	duration  time.Duration
	url       *url.URL
	proxy     *url.URL
	userAgent string
	request   *http.Request
	client    *http.Client
}

func (q *query) Parse(addr string) (*http.Response, error) {
	fmt.Printf("parse addr: %s\n", addr)
	q.setURL(addr)
	q.makeNewClient()
	q.makeNewRequest()
	return q.getResponse()
}

func (q *query) setURL(addr string) {
	fmt.Printf("set url as %s\n", addr)
	parseURL, err := url.Parse(addr)
	if err != nil {
		log.Println(err)
	}
	q.url = parseURL
}

func (q *query) makeNewClient() {
	q.client = &http.Client{
		Timeout: q.timeout * time.Second,
	}
	if q.proxy != nil {
		q.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(q.proxy),
		}
	}
}

func (q *query) makeNewRequest() {
	request, err := http.NewRequest("GET", q.url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	q.request = request
	if q.userAgent != "" {
		q.request.Header.Add("User-Agent", q.userAgent)
	}
}

func (q *query) SetUserAgent(userAgent string) {
	fmt.Printf("set user-agent as %s\n", userAgent)
	q.userAgent = userAgent
}

func (q *query) getResponse() (*http.Response, error) {
	fmt.Printf("get response to request\n")
	response, err := q.client.Do(q.request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (q *query) SetProxy(prx *Proxy) {
	fmt.Printf("set proxy %s:%d\n", prx.ip, prx.port)
	prxParse, err := url.Parse(fmt.Sprintf("tls://%s:%d", prx.ip, prx.port))
	if err != nil {
		log.Println(err)
	}
	q.proxy = prxParse
}

func (q *query) setTimeOut(timeout time.Duration) {
	q.timeout = timeout
}
