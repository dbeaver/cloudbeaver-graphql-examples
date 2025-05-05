package graphql

import (
	"net/http"
	"net/http/cookiejar"
)

func StandardHttpClient() (*http.Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &http.Client{Jar: cookieJar}, nil
}
