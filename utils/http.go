package utils

import (
	"net/http"
	"strings"

	"github.com/altopm/alto/errors"
)

func SendHttpRequest(url string, method string, body string) (string, error) {
	if method == "GET" {
		resp, err := http.Get(url)
		if err != nil {
			errors.Handle(err.Error())
		}
		defer resp.Body.Close()
	}
	if method == "POST" {
		resp, err := http.Post(url, "application/json", strings.NewReader(body))
		if err != nil {
			errors.Handle(err.Error())
		}
		defer resp.Body.Close()
	}
	return "", nil
}
