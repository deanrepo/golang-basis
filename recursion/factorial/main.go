// recurison function
package main

import "fmt"

func fact(n uint) uint {
	if n >= 1 {
		return n * fact(n-1)
	}

	return 1

}

func main() {
	var x uint
	x = 4
	fmt.Println(fact(x))
}
