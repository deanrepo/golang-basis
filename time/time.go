// Time practises
package main

import (
	"fmt"
	"time"
)

// var week time.Duration

func main() {
	// t := time.Now()
	// fmt.Println(t)
	// fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	// t = time.Now().UTC()
	// fmt.Println(t)
	// fmt.Println(time.Now())

	// calculating times:
	// week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	// weekFromNow := t.Add(week)
	// fmt.Println(weekFromNow)

	// formatting times:
	// fmt.Println(t.Format(time.RFC822))
	// fmt.Println(t.Format(time.ANSIC))
	// fmt.Println(t.Format("02 Jan 2006 15:04"))
	// s := t.Format("20060102")
	// fmt.Println(t, "=>", s)

	timeTest1()
}

func timeTest1() {
	t := time.Now().Local()
	td := time.Date(2009, time.February, 23, 0, 0, 0, 0, t.Location())
	td1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	fmt.Println(t)
	fmt.Println(td)
	fmt.Println(td1)
}
