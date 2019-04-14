package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var contentStr = `{
	"name": "thinkingfioa",
	"title": "Go-In-Action",
	"contact": {
		"age":25,
		"tel":"13738142759"
	}
}`

func main() {
	// 使用Map类型

	var c map[string]interface{}
	err := json.Unmarshal([]byte(contentStr), &c)

	if err != nil {
		log.Println("Json Decoder fail")
		return
	}

	fmt.Println(c)
}
