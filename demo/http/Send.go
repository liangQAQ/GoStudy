package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	urlStr := "http://localhost:9090/aa"
	resp, err := http.Get(urlStr)
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	resp1, err1 := http.PostForm(urlStr, url.Values{"key": {"Value"}, "id": {"123"}, "url_long": {"123321"}})
	fmt.Println(resp)
	fmt.Println(err)
	fmt.Println(resp1)
	fmt.Println(err1)
}
