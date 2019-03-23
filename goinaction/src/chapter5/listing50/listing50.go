package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("name %s, email %s\n", u.name, u.email)
}

type admin struct {
	user
	password string
}

func main() {

	ad := admin{
		user: user{
			name:  "thinking_fioa",
			email: "thinking_fioa@163.com",
		},
		password: "123456",
	}
	// 可以直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法也被提升到外部类型中
	ad.notify()
}
