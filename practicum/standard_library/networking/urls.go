package main

import (
	"fmt"
	"net/url"
)

func parseUrl() {
	urlStr := "https://localhost:8000/profile?username=brbryant"

	result, _ := url.Parse(urlStr)

	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // localhost
	fmt.Println(result.Path)     // /profile
	fmt.Println(result.Port())   // 8000
	fmt.Println(result.RawQuery) // username=brbryant

	vals := result.Query()

	fmt.Println(vals["username"])

	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost",
		Path:     "/example",
		RawQuery: "isOnline=true",
	}

	newUrlStr := newUrl.String()

	fmt.Println(newUrlStr)

	newVals := url.Values{}

	newVals.Add("x", "1")
	newVals.Add("y", "2")

	newUrl.RawQuery = newVals.Encode()

	fmt.Println(newUrl.String())
}
