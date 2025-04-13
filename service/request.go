package service

import (
	"io"
	"log"
	"net/http"
	"strings"
)

var client = http.Client{}

// HttpGet get
func HttpGet(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	return httpClient(req)
}

func HttpPost(url string, payload string) string {
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	return httpClient(req)
}

func httpClient(req *http.Request) string {

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	bodyText, _ := io.ReadAll(resp.Body)
	return string(bodyText)
}
