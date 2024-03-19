package util

import (
	"io"
	"net/http"
)

func Fetch(url string, maxRetryTimes int) (string, error) {
	retryTime := 0
	var err error
	for retryTime < maxRetryTimes {
		resp, err := http.Get(url)
		if err != nil {
			retryTime++
			continue
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			retryTime++
			continue
		}
		return string(data), err
	}
	return "", err
}
