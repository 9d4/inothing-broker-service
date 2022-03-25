package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestResponseReadString(t *testing.T) {
	res, err := http.Get("https://google.com")
	if err != nil {
		t.Error(err)
	}

	if len(ResponseReadString(res)) < 0x64 {
		t.Fatal("ResponseReadString should return html of google.com which has length > 100, but got", len(ResponseReadString(res)))
	}
}

func TestGetUrl(t *testing.T) {
	topicApi := getUrl("/api/queues/inothing/traper@air%2fhumid")
	if strings.Contains(topicApi, "/api/queues/inothing/traper@air%252fhumid") {
		t.Fatal("getUrl should only append with host, not changing the query", topicApi)
	}
	fmt.Println(topicApi)
}
