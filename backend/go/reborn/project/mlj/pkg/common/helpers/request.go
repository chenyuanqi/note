package helpers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

// PostForm 发送 HTTP POST 请求
func PostForm(requestURL string, params map[string][]string) ([]byte, error) {
	resp, err := http.PostForm(
		requestURL,
		params,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// PostJSON 请求 JSON
func PostJSON(requestURL string, jsonStr []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// HTTPGet 发送 HTTP GET 请求
func Get(requestURL string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	// 设置 timeout 为5s
	var t int64 = 5
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)

	// 设置参数
	if len(params) > 0 {
		p := req.URL.Query()
		for k, v := range params {
			p.Add(k, v)
		}
		req.URL.RawQuery = p.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
