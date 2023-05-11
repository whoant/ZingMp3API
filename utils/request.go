package utils

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func requestClient() *resty.Client {
	return resty.New().SetRetryCount(3).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() == http.StatusBadGateway || r.StatusCode() == http.StatusServiceUnavailable || r.StatusCode() == http.StatusGatewayTimeout
		})
}

func HttpPost(url string, data interface{}, headers map[string]string) (string, http.Header, error) {
	client := requestClient().R()
	if headers != nil {
		client.SetHeaders(headers)
	}
	if data != nil {
		client.SetBody(data)
	}
	resp, err := client.Post(url)

	if err != nil {
		return "", resp.Header(), fmt.Errorf("faild request : %v", err)
	}
	return string(resp.Body()), resp.Header(), nil
}

func HttpGet(url string, headers map[string]string) (string, http.Header, error) {
	client := requestClient().R()
	if headers != nil {
		client.SetHeaders(headers)
	}
	resp, err := client.Get(url)

	if err != nil {
		return "", resp.Header(), fmt.Errorf("faild request : %v", err)
	}
	return string(resp.Body()), resp.Header(), nil
}
