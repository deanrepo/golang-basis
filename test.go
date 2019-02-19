package main

import "fmt"

func main() {
	var users []string
	// users := []string

	fmt.Println(len(users))
	for i, name := range users {
		fmt.Printf("index: %d, name: %s\n", i, name)
	}
}
