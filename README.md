# Parser web data
А package for simple and convenient collection of information on the Internet.

## Install
```
go get github.com/simpledevdima/pwd 
```

## Example
```go
package main

import (
	"fmt"
	"github.com/simpledevdima/pwd"
	"net/url"
	"strings"
)

func testGet(p *pwd.Http) {
	p.Parse("GET", "https://httpbin.org/get", nil)
	if p.Code != 0 {
		fmt.Println(p.Response.Header)
		fmt.Println(string(p.GetBody()))
	}
}

func testPost(p *pwd.Http) {
	args := make(url.Values)
	args.Set("name", "Dima")
	p.Headers.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	p.Parse("POST", "https://httpbin.org/post", strings.NewReader(args.Encode()))
	if p.Code != 0 {
		fmt.Println(p.Response.Header)
		fmt.Println(string(p.GetBody()))
	}
}

func main() {
	p := pwd.NewHttp(10)
	testGet(p)
	testPost(p)
}
```