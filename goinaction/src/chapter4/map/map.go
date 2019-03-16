package main

import "fmt"

func main() {

	dict := map[string]string{"name": "thinking", "age": "18", "phone": "13738142759"}

	// 第一种方法
	ageValue, exist := dict["age"]
	if exist {
		fmt.Println("age is  %s", ageValue)
	}
	// 第二中方法
	phoneValue := dict["phone"]
	if "" != phoneValue {
		fmt.Println("phone is %d", phoneValue)
	}

	// 遍历
	for key, value := range dict {
		fmt.Println("key is %s, value is %s", key, value)
	}

	delete(dict, "age")

	// 遍历
	for key, value := range dict {
		fmt.Println("key is %s, value is %s", key, value)
	}
}
