package main

import (
	"fmt"
	"os"
)

func main() {
	// var w io.Writer
	// w = os.Stdout
	// w.Write([]byte("hello world"))
	// fmt.Printf("\n%T\t%+v\n", w, w)
	// f := w.(*os.File) // success: f == os.Stdout
	// fmt.Printf("%T\t%+v\n", f, f)

	// rw := w.(io.ReadWriter)
	// fmt.Printf("\n%T\t%+v\n", rw, rw)

	// c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	// fmt.Printf("%T\t%+v\n", c, c)

	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)

}
