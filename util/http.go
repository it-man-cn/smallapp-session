package util

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (r []byte, err error) {
	var result []byte
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		result = body
		return result, nil
	}
	return result, err
}
