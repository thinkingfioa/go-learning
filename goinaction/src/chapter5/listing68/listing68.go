package main

import (
	"./counters"
	"fmt"
)

func main() {
	counter := counters.New(1)
	fmt.Printf("count %d", counter)
}
