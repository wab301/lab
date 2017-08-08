package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 使用代理
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}
	tr := &http.Transport{
		Proxy: proxy,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:9090/hello")
	// 修改https证书验证，不再对服务端的证书进行校验
	// resp, err := http.Get("https://localhost:9090")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("=====", string(body))
}
