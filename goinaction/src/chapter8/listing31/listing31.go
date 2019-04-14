package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})

	c["name"] = "thinkingfioa"
	c["title"] = "Go-In-Action"
	c["contact"] = map[string]interface{}{
		"age": 15,
		"tel": "13838383838",
	}

	//data, err := json.MarshalIndent(c, "", "    ")
	data, err := json.Marshal(c)

	if err != nil {
		log.Println("Json Encoder fail.")
		return
	}

	fmt.Println(string(data))
}
