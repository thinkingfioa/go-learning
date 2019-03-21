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

func main() {
	u := user{
		name: "thinking", email: "thinking_fioa@163.com",
	}

	sendNotification(u)

	uPointer := &user{
		name: "ppp", email: "ppp@163.com",
	}

	sendNotification(uPointer)
}

func sendNotification(n notifier) {
	n.notify()
}
