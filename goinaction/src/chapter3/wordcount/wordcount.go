package main

import (
	"../words"
	"fmt"
	"io/ioutil"
)

const filename = "./goinaction/src/chapter3/wordcount/gowords.txt"

func main() {
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Print(err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)

	fmt.Printf("Thread have %d words", count)
}
