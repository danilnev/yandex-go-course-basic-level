package main

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write(p []byte) int
}

type Reader interface {
	Read() []byte
}

type ReaderWriter interface {
	Reader
	Writer
}

type UpperReaderWriter struct {
	UpperString string
}

func (upw *UpperReaderWriter) Write(p []byte) int {
	upw.UpperString = strings.ToUpper(string(p))
	return len(p)
}

func (upw *UpperReaderWriter) Read() []byte {
	return []byte(upw.UpperString)
}

func main() {
	var rw ReaderWriter = &UpperReaderWriter{}

	inputString := "Hello, World!"
	len_upw := rw.Write([]byte(inputString))

	outputBytes := rw.Read()
	fmt.Println(string(outputBytes) + "\n" + fmt.Sprintf("Length: %d", len_upw)) // Output: "HELLO, WORLD!"
}
