package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// fileWrite()
	// fileRead()
	// bufferedWrite()
	// bufferedRead()
	appendFile()
}

// fileWrite write something to a file.
func fileWrite() {
	file, err := os.Create("myfile.txt")
	if err != nil {
		panic(err)
	}

	data := []byte("Hello world!\n")
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	file.WriteString("This is a test.")
	file.Close()
}

// fileRead reads something from a file.
func fileRead() {
	// Open the file.
	file, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 5)
	file.Seek(-6, 2)
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	file.Seek(0, 0)
	newData := make([]byte, 5)
	_, err = file.Read(newData)
	fmt.Println(string(newData))
}

// bufferedWrite bufferd write something to a file.
func bufferedWrite() {
	file, err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffer for the os file
	buf := bufio.NewWriter(file)
	// Write data to buf
	buf.WriteString("Buffered write to a file.")
	// Clears the buf data and writes to the file.
	err = buf.Flush()
	if err != nil {
		log.Fatal(err)
	}

}

// bufferedRead buffered reads from the file.
func bufferedRead() {
	file, err := os.Open("newFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := reader.Peek(5)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

// appendFile appends content to the file.
func appendFile() {
	file, err := os.OpenFile("newFile.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("trying to append.\n")
}
