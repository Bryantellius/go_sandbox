package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

const url = "https://raw.githubusercontent.com/Bryantellius/kingkillerapi/main/assets/terms.json"

func main() {
	response, err := http.Get(url)

	checkError(err)

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)

	checkError(err)

	var terms []Term
	err = json.Unmarshal(bytes, &terms)
	checkError(err)

	for _, term := range terms {
		fmt.Println(term.Name)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Term struct {
	Name string `json:"term"`
	Def string `json:"def"`
	Src string `json:"src"`
}
