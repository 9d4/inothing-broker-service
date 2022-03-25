package main

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Convenient wrapper to make request to RabbitMQ through http.
// This wrapper sets auth automatically.
func Req(method string, path string, body io.Reader) (*http.Response, error) {
	client := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(method, getUrl(path), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(RABBITMQ_USER, RABBITMQ_PASS)

	return client.Do(req)
}

// Read res.body and return as string.
// if any error occurs will return ""
func ResponseReadString(res *http.Response) string {
	if res == nil {
		return ""
	}

	lines, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return ""
	}

	return string(lines)
}

func getUrl(p string) string {
	u, err := url.Parse(RABBITMQ_HOST)
	if err != nil {
		return ""
	}

	u.Path = path.Join(u.Path, p)
	final, err := url.PathUnescape(u.String())
	if err != nil {
		return RABBITMQ_HOST
	}

	return final
}
