package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("name %s, email %s.", u.name, u.email)
	fmt.Println()
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	u := user{"thinking", "thinking_fioa@163.com"}

	u.notify()
	u.changeEmail("thinkingfioa@163.com")
	u.notify()

	uPoint := &user{"triple3P", "triple3P@163.com"}
	uPoint.notify()
	uPoint.changeEmail("sb@163.com")
	uPoint.notify()
}
