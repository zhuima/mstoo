package pkg

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

const (
	HTTP_MIN_STATUS = 200
	HTTP_MAX_STATUS = 299
)

type Link struct {
	Url    *url.URL `json:"url"`
	Status int      `json:"status"`
}

func NewHttpRequest() *http.Client {
	transport := &http.Transport{
		MaxIdleConnsPerHost: 100,  // default = 2
		MaxIdleConns:        6000, // default = 100
		IdleConnTimeout:     30 * time.Second,
		DisableCompression:  true,
		MaxConnsPerHost:     500, // default = 0

		// ResponseHeaderTimeout: 20 * time.Second,
		// TLSHandshakeTimeout:   10 * time.Second,
		// ExpectContinueTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: func(req *http.Request) (*url.URL, error) {
			// req.Header.Add("", "")
			return nil, nil
		},

		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}
	return client
}

func (l *Link) IsHealthy() bool {
	return l.Status >= HTTP_MIN_STATUS && l.Status <= HTTP_MAX_STATUS
}
