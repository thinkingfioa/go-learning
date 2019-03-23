package main

import (
	"./entities"
	"fmt"
)

func main() {

	ad := entities.Admin{
		Passwd: "123456",
	}
	ad.Name = "thinking"
	ad.Email = "thinking_fiao@163.com"

	fmt.Printf("name %s, email %s, passwd %s", ad.Name, ad.Email, ad.Passwd)
}
