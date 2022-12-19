package pwd

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"time"
)

// NewHttp creates a new Http parser and returns a link to it
func NewHttp(timeout time.Duration) *Http {
	h := &Http{}
	h.Headers = &http.Header{}
	h.SetTimeOut(timeout)
	return h
}

// Http data type containing basic query settings
type Http struct {
	timeout  time.Duration
	duration time.Duration
	Headers  *http.Header
	url      *url.URL
	method   string
	body     io.Reader
	proxy    *Proxy
	request  *http.Request
	client   *http.Client
	Response *http.Response
	Code     int
}

// SetTimeOut sets the value of the timeout parameter
func (h *Http) SetTimeOut(timeout time.Duration) {
	h.timeout = timeout
}

// SetProxy sets the proxy server with which the request will be made
func (h *Http) SetProxy(prx *Proxy) {
	fmt.Printf("set proxy id=%d addr=%s:%d\n", prx.id, prx.ip, prx.port)
	h.proxy = prx
}

// setUrl sets the value of the url parameter
func (h *Http) setURL(link string) {
	fmt.Printf("set url as %s\n", link)
	parseURL, err := url.Parse(link)
	if err != nil {
		log.Println(err)
	}
	h.url = parseURL
}

// setMethod sets the value of the request method
func (h *Http) setMethod(method string) {
	fmt.Printf("set method as %s\n", method)
	h.method = method
}

// setBody sets the value of the request body
func (h *Http) setBody(body io.Reader) {
	fmt.Printf("set body as %s\n", body)
	h.body = body
}

// makeClient creating a client that will make a request
func (h *Http) makeClient() {
	h.client = &http.Client{
		Timeout: h.timeout * time.Second,
	}
	if h.proxy != nil {
		// set proxy to client
		prxParse, err := url.Parse(fmt.Sprintf("tls://%s:%d", h.proxy.ip, h.proxy.port))
		if err != nil {
			log.Println(err)
		}
		h.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(prxParse),
		}
	}
}

// makeRequest creating a request
func (h *Http) makeRequest() {
	request, err := http.NewRequest(h.method, h.url.String(), h.body)
	if err != nil {
		log.Println(err)
	}
	h.request = request
	h.request.Header = *h.Headers // add request headers
}

// doRequest executing a request and returning a response and an error value
func (h *Http) doRequest() {
	fmt.Printf("get response to request\n")
	startRequest := time.Now()
	response, err := h.client.Do(h.request)
	durationRequest := math.Round(time.Since(startRequest).Seconds()*100) / 100
	if err != nil {
		log.Println(err)
	}
	h.Response = response
	if h.Response != nil {
		h.Code = h.Response.StatusCode
	} else {
		h.Code = 0
	}
	if h.proxy != nil {
		// logging proxy
		h.proxy.Log(h.url.String(), h.Code, durationRequest)
	}
}

// GetBody returns the body of the response to the completed request
func (h *Http) GetBody() []byte {
	body, err := io.ReadAll(h.Response.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}

// Parse receiving a response from the server in accordance with the request parameters
func (h *Http) Parse(method string, link string, body io.Reader) {
	fmt.Printf("parse addr: %s\n", link)
	h.setMethod(method)
	h.setURL(link)
	h.setBody(body)
	h.makeClient()
	h.makeRequest()
	h.doRequest()
}

// GetRandUserAgent returns a random User-Agent from the database
func (h *Http) GetRandUserAgent(db *sql.DB) string {
	var name string
	err := db.QueryRow("select `name` from `user_agents` order by rand() limit 1").Scan(&name)
	if err != nil {
		log.Println(err)
	}
	return name
}
