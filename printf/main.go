// Printf format test
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	// p := point{1, 2}
	// fmt.Printf("%v\n", p)

	// fmt.Printf("%+v\n", p)

	// fmt.Printf("%#v\n", p)

	// fmt.Printf("%T\n", p)

	// fmt.Printf("%t\n", true)

	// fmt.Printf("%d\n", 123)

	// fmt.Printf("%b\n", 14)

	// fmt.Printf("%c\n", 33)

	// fmt.Printf("%x\n", 456)

	// fmt.Printf("%f\n", 78.9)

	// fmt.Printf("%e\n", 12340000.0)
	// fmt.Printf("%E\n", 12340000.0)

	// fmt.Printf("%s\n", "\"string\"")
	// fmt.Printf("%9.9s\n", "\"string\"")
	// fmt.Printf("%.9s\n", "\"string\"")

	// fmt.Printf("%q\n", "\"string\"")

	// fmt.Printf("%x\n", "hex this")

	// fmt.Printf("%p\n", &p)

	// fmt.Printf("|%6d|%6d|\n", 12, 345)

	// fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// s := fmt.Sprintf("a %s", "string")
	// fmt.Println(s)

	s1 := "this is just a test"
	s2 := "this ia another test"

	fmt.Printf("%s %s\n", s1, s2)

	// %9.9s=>width 9, show the first 9 characters of the string; %.55=>show the first 55 chatacters of the string
	fmt.Printf("%9.9s %.55s\n", s1, s2)
	fmt.Printf("%10.9s %.55s\n", s1, s2)

}
