package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Action struct {
	Payload interface{} `json:"payload"`
	Type    string      `json:"type"`
}

func main() {
	postRequestTest()
}

func postRequestTest() {
	const httpbinUrl = "https://httpbin.org/post"

	testPayload := Action{
		Payload: map[string]string{
			"name": "Frodo",
			"age":  "51",
		},
		Type: "NEW_USER",
	}

	testPayloadJson, err := json.Marshal(testPayload)
	handleErr(err)

	// Sends POST request to url
	response, err := http.Post(httpbinUrl, "application/json", bytes.NewReader(testPayloadJson))
	handleErr(err)

	content, err := io.ReadAll(response.Body)

	handleErr(err)

	var contentJson map[string]interface{}

	json.Unmarshal(content, &contentJson)

	defer response.Body.Close()

	fmt.Printf("%v\n", contentJson)

	fmt.Println(contentJson["url"])
}

func getRequestTest() {
	const httpbinUrl = "https://httpbin.org/get?a=1&b=2"

	// Sends GET request to url
	response, err := http.Get(httpbinUrl)

	handleErr(err)

	// Closes the response stream
	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Protocol:", response.Proto)
	fmt.Println("Content Length:", response.ContentLength)

	var sb strings.Builder

	content, err := io.ReadAll(response.Body)
	handleErr(err)

	byteCount, _ := sb.Write(content)
	handleErr(err)

	fmt.Println(byteCount, sb.String())
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
