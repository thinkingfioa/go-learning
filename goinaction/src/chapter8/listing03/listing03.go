package main

import "log"

func init() {
	log.SetPrefix("TRICE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println写到标准日志记录器
	log.Println("message")

	// Fatalln调用Println()函数后，调用os.Exit(1)。所以下面代码不会执行
	log.Fatalf("fatal message")

	// Panicln调用Println()函数后，调用panic()
	log.Panicln("panic message")
}
