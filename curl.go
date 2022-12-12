package pwd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func NewCurl(timeout time.Duration) *Curl {
	c := &Curl{}
	c.setTimeOut(timeout)
	return c
}

type Curl struct {
	timeout   time.Duration
	duration  time.Duration
	url       *url.URL
	proxy     *url.URL
	userAgent string
	request   *http.Request
	client    *http.Client
}

func (c *Curl) Parse(addr string) (*http.Response, error) {
	fmt.Printf("parse addr: %s\n", addr)
	c.setURL(addr)
	c.makeNewClient()
	c.makeNewRequest()
	return c.getResponse()
}

func (c *Curl) setURL(addr string) {
	fmt.Printf("set url as %s\n", addr)
	parseURL, err := url.Parse(addr)
	if err != nil {
		log.Println(err)
	}
	c.url = parseURL
}

func (c *Curl) makeNewClient() {
	c.client = &http.Client{
		Timeout: c.timeout * time.Second,
	}
	if c.proxy != nil {
		c.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(c.proxy),
		}
	}
}

func (c *Curl) makeNewRequest() {
	request, err := http.NewRequest("GET", c.url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	c.request = request
	if c.userAgent != "" {
		c.request.Header.Add("User-Agent", c.userAgent)
	}
}

func (c *Curl) SetUserAgent(userAgent string) {
	fmt.Printf("set user-agent as %s\n", userAgent)
	c.userAgent = userAgent
}

func (c *Curl) getResponse() (*http.Response, error) {
	fmt.Printf("get response to request\n")
	response, err := c.client.Do(c.request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Curl) SetProxy(prx *Proxy) {
	fmt.Printf("set proxy %s:%d\n", prx.ip, prx.port)
	prxParse, err := url.Parse(fmt.Sprintf("tls://%s:%d", prx.ip, prx.port))
	if err != nil {
		log.Println(err)
	}
	c.proxy = prxParse
}

func (c *Curl) setTimeOut(timeout time.Duration) {
	c.timeout = timeout
}
