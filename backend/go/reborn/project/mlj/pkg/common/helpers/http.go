package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type httpSrc struct {
	Request  *http.Request
	Response *http.Response
}

func HTTP() *httpSrc {
	return &httpSrc{}
}

func (h *httpSrc) Get(uri string, headSet ...map[string]string) (string, error) {
	req, _ := http.NewRequest("GET", uri, nil)
	var header map[string]string
	if len(headSet) > 0 {
		header = headSet[0]
	}
	return h.do(req, header)
}

func (h *httpSrc) Post(uri string, data map[string]string, headSet ...map[string]string) (string, error) {
	values := url.Values{}
	for key, val := range data {
		values.Add(key, val)
	}
	req, _ := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	var header map[string]string
	if len(headSet) > 0 {
		header = headSet[0]
	}
	return h.do(req, header)
}

func (h *httpSrc) do(req *http.Request, header map[string]string) (string, error) {
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:61.0) Gecko/20100101 Firefox/61.0")
	client := &http.Client{}
	if len(header) > 0 {
		// 设置cookie
		if setCookie, ok := header["cookie"]; ok {
			var cookies []*http.Cookie
			if err := json.Unmarshal([]byte(setCookie), &cookies); err != nil {
				return "", errors.New(fmt.Sprintf("set cookies error: %s", err.Error()))
			}
			for _, cookie := range cookies {
				req.AddCookie(cookie)
			}
			delete(header, "cookie")
		}
		// 重定向次数修改
		if checkNum, ok := header["CheckRedirect"]; ok {
			cNum, _ := strconv.Atoi(checkNum)
			if cNum != 10 {
				client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
					if cNum == 0 || len(via) < cNum {
						return nil
					}
					return fmt.Errorf("stopped after %d redirects", cNum)
				}
			}
		}
		// 设置头部信息
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	h.Request = req
	res, err := client.Do(req)
	h.Response = res
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

func ParseResp(resp *http.Response, out interface{}) (body []byte, err error) {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("readall err: %s", err)
	}

	if bytes.Contains(b, []byte(`"result":"error"`)) {
		return b, fmt.Errorf("result err: %s", b)
	}

	if out != nil {
		if err = json.Unmarshal(b, out); err != nil {
			return b, fmt.Errorf("body decode err: %s", err)
		}
	}

	return b, nil
}

func GetClientIP(req *http.Request) string {
	reward := req.Header.Get("X-Forwarded-For")
	if reward != "" {
		reward = strings.ReplaceAll(reward, " ", "")
		return strings.Split(reward, ",")[0]
	}
	if req.RemoteAddr != "" {
		if strings.Contains(req.RemoteAddr, "[::1]") {
			return "127.0.0.1"
		}
		return strings.Split(req.RemoteAddr, ":")[0]
	}
	return ""
}

func ApiRetry(f func() (interface{}, error), retry int, delay time.Duration) (res interface{}, err error) {
	if delay == 0 {
		delay = 100 * time.Millisecond
	}
	for i := 1; i <= retry+1; i++ {
		res, err = f()
		if err == nil {
			return
		}
		if delay > 0 {
			time.Sleep(delay)
		}
	}

	return nil, fmt.Errorf("api retry max times, still err: %s", err)
}
