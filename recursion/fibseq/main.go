package main

import "fmt"

func fib(n uint) uint {
	if n >= 3 {
		return fib(n-1) + fib(n-2)
	}

	return 1
}

func main() {
	var x uint
	x = 6
	fmt.Println(fib(x))
}
