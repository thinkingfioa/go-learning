package main

import (
	"fmt"
	"net"
)

func CheckError2(err error) {
	if nil != err {
		fmt.Println("Fatal error :%s", err.Error())
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9291")
	CheckError2(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError2(err)

	fmt.Println("connection success. %s", conn.RemoteAddr().String())

	sender(conn)
}

func sender(conn net.Conn) {
	words := "hello word"

	conn.Write([]byte(words))

	// 接收数据
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	CheckError2(err)

	fmt.Println("receive server back msg. %s", buffer[:n])
}
