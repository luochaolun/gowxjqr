package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	//"compress/gzip"
)

func GET(targetUrl string) (bool, string) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", targetUrl, nil) //建立一个请求
	if err != nil {
		return false, ""
	}
	//Add 头协议
	reqest.Header.Set("Accept", "text/xml")
	//reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	//reqest.Header.Del("Accept-Encoding")
	reqest.Header.Set("Accept-Encoding", "gzip")
	//reqest.Header.Add("Connection", "keep-alive")
	//reqest.Header.Add("Cookie", "设置cookie")
	reqest.Header.Add("Referer", "http://dns.weixin.qq.com/")
	//reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest) //提交
	defer response.Body.Close()

	cookies := response.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, ""
	}
	/*reader, _ := gzip.NewReader(response.Body)
	var body string
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)

		if err != nil && err != io.EOF {
			 panic(err)
		}

		if n == 0 {
			 break
		}
		body += string(buf)
	}*/

	return true, string(body)
}

func main() {
	bRet, html := GET("http://dns.weixin.qq.com/cgi-bin/micromsg-bin/newgetdns")
	if !bRet {
		fmt.Println("访问出错！")
		return
	}
	fmt.Println(html)
}
