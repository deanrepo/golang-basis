package message

import "fmt"

var Msg Message = construct()

func init() {
	fmt.Print("messsage pkg init function run, print->")
	Msg.PrintMsg()
	fmt.Println()
}

type Message struct {
	Key   string
	Value string
}

func (msg Message) PrintMsg() {
	fmt.Printf("message: %+v\n", msg)
}

func (msg Message) Test() {
	fmt.Println("This is a test!")
}

func construct() Message {
	fmt.Printf("message pkg construct run\n\n")
	return Message{
		"party",
		"shit",
	}
}
