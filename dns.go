package main

import (
	"compress/flate"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

// 解码
func Gzdecode(data string) string {
	if data == "" {
		return ""
	}
	r := flate.NewReader(strings.NewReader(data))
	defer r.Close()
	out, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Errorf("%s\n", err)
		return ""
	}
	return string(out)
}

func GET(targetUrl string) (bool, string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", targetUrl, nil) //建立一个请求
	if err != nil {
		return false, ""
	}

	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	resp, err := client.Do(req) //提交
	defer resp.Body.Close()

	//fmt.Println(resp.Header)
	/*cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}*/

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, ""
	}

	html := string(body)
	if resp.Header.Get("Content-Encoding") == "deflate" {
		html = Gzdecode(html)
	}

	return true, html
}

func getDns() (bool, []string, []string) {
	sDns := []string{}
	lDns := []string{}

	bRet, html := GET("http://dns.weixin.qq.com/cgi-bin/micromsg-bin/newgetdns")
	if !bRet {
		return false, sDns, lDns
	}
	html = strings.Replace(html, "domain", "div", -1)
	html = strings.Replace(html, "name=", "class=", -1)
	html = strings.Replace(html, ".weixin.qq.com", "weixinqqcom", -1)
	//fmt.Println(html)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return false, sDns, lDns
	}

	doc.Find("div.shortweixinqqcom ip").Each(func(_ int, s *goquery.Selection) {
		sDns = append(sDns, s.Text())
	})

	doc.Find("div.longweixinqqcom ip").Each(func(_ int, s *goquery.Selection) {
		lDns = append(lDns, s.Text())
	})

	if len(sDns) == 0 || len(lDns) == 0 {
		return false, sDns, lDns
	}

	return true, sDns, lDns
}
