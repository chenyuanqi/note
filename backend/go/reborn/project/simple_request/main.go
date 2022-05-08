package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// get 请求
	getRsps, err := http.Get("https://v0.yiketianqi.com/api?unescape=1&version=v61&appid=85841439&appsecret=EKCDLT4I")
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer getRsps.Body.Close()

	getBody, err := ioutil.ReadAll(getRsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}

	fmt.Println(string(getBody))

	// post 请求
	postRsps, err := http.Post("http://example.com/", "plain/text", nil)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer postRsps.Body.Close()

	postBody, err := ioutil.ReadAll(postRsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}

	fmt.Println(string(postBody))

	// 设置超时
	var client = &http.Client{
		Timeout: time.Second * 5, // 超时设置
	}

	rqst, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Println("New request failed:", err)
		return
	}

	// 设置头部
	rqst.Header.Add("X-My-Auth", "xxxx")

	rsps, err := client.Do(rqst)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}

	fmt.Println(string(body))
}
