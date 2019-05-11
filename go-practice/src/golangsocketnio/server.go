package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if nil != err {
		fmt.Println("Fatal error :%s", err.Error())
	}
}

func main() {
	netListen, err := net.Listen("tcp", "localhost:9291")
	CheckError(err)

	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		CheckError(err)

		fmt.Println("accept conn %s success", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	// TODO 粘包处理
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		CheckError(err)

		fmt.Println("receive data str: %s", string(buffer[:n]))
		conn.Write(buffer[:n])
	}
}
