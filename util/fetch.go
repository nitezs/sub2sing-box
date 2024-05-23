package util

import (
	"io"
	"net/http"
)

func Fetch(url string, maxRetryTimes int) (string, error) {
	retryTime := 0
	var err error
	var resp *http.Response
	for retryTime < maxRetryTimes {
		resp, err = http.Get(url)
		if err != nil {
			retryTime++
			continue
		}
		var data []byte
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			retryTime++
			continue
		}
		return string(data), err
	}
	return "", err
}
