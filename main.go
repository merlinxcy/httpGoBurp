package main

import (
	"fmt"
	"httphacker/httpHack"
)

func main(){
	raw := `
GET /get?a=1&b=2&c=heloo HTTP/1.1
Host: httpbin.org
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: celebrate hack-requests 1.0 !
Accept: te:sssxt/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
asdsd:as:das
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
Cookie: _gauges_unique_hour=1; _gauges_unique_day=1; _gauges_unique_month=1; _gauges_unique_year=1; _gauges_unique=1

sd
`
	raw = `
POST /post HTTP/1.1
Host: httpbin.org
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: celebrate hack-requests 1.0 !
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
Cookie: _gauges_unique_hour=1; _gauges_unique_day=1; _gauges_unique_month=1; _gauges_unique_year=1; _gauges_unique=1

a=1&b=2&c=heloo
`
	http := new(httpHack.Http)
	fmt.Println(http.DoHttpRaw(raw))
}
