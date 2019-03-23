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
	fmt.Printf("name is %s, email is %s\n", u.name, u.email)
}

type admin struct {
	user
	password string
}

func (ad admin) notify() {
	fmt.Printf("name is %s, email is %s, passwd is %s\n", ad.name, ad.email, ad.password)
}

func main() {
	ad := admin{
		user: user{
			name:  "thinking",
			email: "thinking_fioa@163.com",
		},
		password: "123456",
	}

	sendNotification(ad.user)
	sendNotification(ad)
}

func sendNotification(n notifier) {
	n.notify()
}
