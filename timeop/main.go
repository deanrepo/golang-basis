package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// rocketLaunch1()
	rocketLaunch2()
}

func rocketLaunch1() {

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}

	fmt.Println("launch succefully!")

}

func rocketLaunch2() {

	fmt.Println("Commencing countdown. Press return to abort.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read a single byte
		abort <- struct{}{}
	}()

	tick := time.Tick(time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}

	fmt.Println("launch succefully!")
}
