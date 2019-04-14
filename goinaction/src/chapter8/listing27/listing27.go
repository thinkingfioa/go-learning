package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Age int    `json:"age"`
		Tel string `json:"tel"`
	}
}

var contentStr = `{
	"name": "thinkingfioa",
	"title": "Go-In-Action",
	"contact": {
		"age":25,
		"tel":"13738142759"
	}
}`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(contentStr), &c)

	if err != nil {
		log.Println("Json Decoder fail")
		return
	}

	fmt.Println(c)
}
