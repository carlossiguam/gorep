package main

import (
	"bytes"
	"fmt"
	"log"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

func main() {
	method1()
	initApplog()
}

func method1() {
	fmt.Printf("Has ingresado a mi programa a ver de a como nos toca \n")
}

func initApplog() {
	logger.Print("Hello, log file!")
	fmt.Println(logger)
}
