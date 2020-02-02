package httpHack

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Http struct {
	url string
	action string
	header map[string]string
	body string
}

func (d *Http)DoHttpRaw(raw string) string{
	raw = strings.TrimSpace(raw)
	rawList := strings.Split(raw,"\n")
	//http action
	httpFirstLine := rawList[0]
	ret := strings.Split(httpFirstLine," ")
	action := ret[0]
	path := ret[1]
	//protocol := ret[2]
	//http next
	httpNext := rawList[1:]
	bodyPos := 0
	d.header = make(map[string]string)
	for i,j := range httpNext{
		if j == "" {
			bodyPos = i + 1
			break
		}
		ret := strings.Split(j,": ")
		headerName := ret[0]
		headerVaule := "" //fix
		for _,tmp := range ret[1:]{
			headerVaule = headerVaule + tmp
		}
		d.header[headerName] = headerVaule
	}

	var httpBody []string
	if bodyPos != 0 {
		httpBody = rawList[bodyPos:]
	}
	d.body = ""
	for _,j := range httpBody{
		d.body = d.body + j
	}
	d.url = "http://" + d.header["Host"] + path
	fmt.Println(d.url)
	if action == "GET" {
		httpHan,_ := http.NewRequest("GET", d.url, strings.NewReader(raw))
		client := new(http.Client)
		resp,_ :=client.Do(httpHan)
		resp_raw, _ := httputil.DumpResponse(resp,true)
		return string(resp_raw)
	} else if action == "POST" {
		httpHan, _ := http.NewRequest("POST", d.url, strings.NewReader(raw))
		client := new(http.Client)
		resp, _ := client.Do(httpHan)
		resp_raw, _ := httputil.DumpResponse(resp, true)
		return string(resp_raw)
	} else {
		return ""
	}
}