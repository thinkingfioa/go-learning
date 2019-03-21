package main

import "fmt"

type notifierPointer interface {
	notifyPoint()
}

type userPointer struct {
	name  string
	email string
}

func (uPointer *userPointer) notifyPoint() {
	fmt.Print("name %s, email %s", uPointer.name, uPointer.email)
}

func main() {

	uPointer := userPointer{
		name:  "thinking",
		email: "thinking_fioa@163.com",
	}

	sendNotificationPoint(uPointer)

	sendNotificationPoint(&uPointer)
}

func sendNotificationPoint(n notifierPointer) {
	n.notifyPoint()
}
