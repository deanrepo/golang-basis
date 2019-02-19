package main

import (
	"common_files/initfunc/message"
	"fmt"
)

var _ int64 = s()

func init() {
	fmt.Println("init in sandbox.go")
}

func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}

func main() {
	fmt.Println("main")
	fmt.Println()

	fmt.Println("calling Msg.Test()")
	message.Msg.Test()
}
